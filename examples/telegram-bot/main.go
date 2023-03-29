package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lodthe/goldb/db"
	"go.uber.org/zap"
)

type Repository struct {
	conn *db.Connection
}

type UserMessage struct {
	Username string
	Message  string
}

func (c *Repository) AddMessage(ctx context.Context, username string, message string) error {
	data, err := json.Marshal(UserMessage{Username: username, Message: message})
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	_, err = c.conn.Put(ctx, "messages", string(data))

	return err
}

// GetAllMessages fetches all messages from the database.
func (c *Repository) GetAllMessages(ctx context.Context) ([]UserMessage, error) {
	iterator, err := c.conn.GetIterator(ctx, db.IterKeyEquals("messages"))
	if err != nil {
		return nil, fmt.Errorf("failed to create iterator: %w", err)
	}

	var messages []UserMessage
	for iterator.HasNext() {
		item, err := iterator.GetNext()
		if err != nil {
			return nil, fmt.Errorf("failed to get next message: %w", err)
		}

		var msg UserMessage
		err = json.Unmarshal([]byte(item.Value), &msg)
		if err != nil {
			// Handle error.
			continue
		}

		messages = append(messages, msg)
	}

	return messages, nil
}

func main() {
	logger, _ := zap.NewDevelopment()

	// Establish a connection with the server.
	conn, err := db.Open(
		db.WithLogger(logger),
		// Provide server address.
		db.WithServerAddress("bloom.lodthe.me:8888"),
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	defer conn.Close()

	repository := &Repository{
		conn: conn,
	}

	token := os.Getenv("TELEGRAM_APITOKEN")
	handleUpdates(token, repository, logger)
}

func handleUpdates(token string, repository *Repository, logger *zap.Logger) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		logger.Fatal("failed to create botapi", zap.Error(err))
	}

	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	// Handle Telegram updates.
	for update := range updates {
		if update.Message == nil {
			continue
		}

		username := update.Message.From.UserName
		userMsg := update.Message.Text

		responseText := "Messages from other users:\n"
		if len(userMsg) > 100 {
			responseText = "Message length cannot exceed 100 symbols"
		} else {
			err := repository.AddMessage(context.Background(), username, userMsg)
			if err != nil {
				logger.Error("failed to save a message:", zap.Error(err), zap.String("message", userMsg))
				continue
			}

			messages, err := repository.GetAllMessages(context.Background())
			if err != nil {
				logger.Error("failed to fetch messages:", zap.Error(err))
				continue
			}

			const suffixLen = 20
			if len(messages) > suffixLen {
				messages = messages[len(messages)-suffixLen:]
			}

			for _, m := range messages {
				responseText += fmt.Sprintf("\n@%s: %s", m.Username, html.EscapeString(m.Message))
			}
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "html"

		_, err = bot.Send(msg)
		if err != nil {
			logger.Error("failed to send a message:", zap.Error(err))
		}
	}
}
