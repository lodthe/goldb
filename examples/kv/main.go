package main

import (
	"context"
	"fmt"

	"github.com/lodthe/goldb/db"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	logger, _ := zap.NewDevelopment()

	conn, err := db.Open(
		db.WithLogger(logger),
		db.WithServerAddress("bloom.lodthe.me:8888"),
	)
	if err != nil {
		logger.Error(err.Error())
	}

	defer conn.Close()

	var key string
	var value string
	_, _ = fmt.Scanf("%s%s", &key, &value)

	triplet, err := conn.Put(ctx, key, value)
	if err != nil {
		logger.Fatal("failed to put", zap.Error(err))
	}

	triplet, err = conn.GetLatest(ctx, key)
	if err != nil {
		logger.Fatal("failed to get latest", zap.Error(err))
	}

	logger.Sugar().Infof("got values: %s -> %s (%s)", triplet.Key, triplet.Value, triplet.Version)
}
