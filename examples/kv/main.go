package main

import (
	"github.com/lodthe/goldb/db"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	conn, err := db.Open(
		db.WithLogger(logger),
		db.WithServerAddress("localhost:13337"),
	)
	if err != nil {
		logger.Error(err.Error())
	}

	defer conn.Close()
}
