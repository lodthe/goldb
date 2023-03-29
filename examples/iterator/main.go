package main

import (
	"context"

	"github.com/lodthe/goldb/db"
	"go.uber.org/zap"
)

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

	options := []db.IterOption{
		db.IterKeyEquals("Alice"),
	}

	iterator, err := conn.GetIterator(context.Background(), options...)
	if err != nil {
		logger.Fatal(err.Error())
	}

	for iterator.HasNext() {
		item, err := iterator.GetNext()
		if err != nil {

		}

		logger.Sugar().Infof("[%s] %s -> %s", item.Version, item.Key, item.Value)
	}
}
