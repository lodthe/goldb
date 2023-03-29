package main

import (
	"context"

	"github.com/lodthe/goldb/db"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
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

	key := "Alice"
	value := "Alice's shopping cart"

	// Create a new record.
	triplet, err := conn.Put(ctx, key, value)
	if err != nil {
		logger.Fatal("failed to put", zap.Error(err))
	}

	// Get the latest value for "Alice" key.
	triplet, err = conn.GetLatest(ctx, key)
	if err != nil {
		logger.Fatal("failed to get latest", zap.Error(err))
	}

	logger.Sugar().Infof("got values: %s -> %s (%s)", triplet.Key, triplet.Value, triplet.Version)
}
