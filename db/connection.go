package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/lodthe/goldb/db/dbclient"
	"go.uber.org/zap"
)

type Connection struct {
	ctx    context.Context
	logger *zap.Logger
	client dbclient.Client
}

// Open creates a new database connection.
// It takes a list of options to customize settings. See config_options.go for more details.
func Open(options ...Option) (conn *Connection, err error) {
	config := defaultConfig

	for _, apply := range options {
		apply(&config)
	}

	// Create gRPC client if necessary.
	client := config.client
	if client == nil {
		grpcConn := config.grpcConn
		if grpcConn == nil {
			grpcConn, err = initGRPCConnection(config.serverAddress)
			if err != nil {
				return nil, fmt.Errorf("failed to establish connection: %w", err)
			}

			// Do not forget to close connection in case of panic.
			defer func() {
				if err != nil {
					grpcConn.Close()
				}
			}()
		}

		client = dbclient.NewGRPCClient(grpcConn)
	}

	return &Connection{
		ctx:    config.ctx,
		logger: config.logger,
		client: client,
	}, nil
}

// Close closes internal connection with the server and cleans resources.
// It only returns nil in the current implementation.
//
// The connection cannot be used after Close was called.
func (c *Connection) Close() error {
	if c.client == nil {
		return errors.New("client hasn't been created")
	}

	err := c.client.CloseConnection()
	if err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}

	c.logger.Info("connection with the database server has been closed")

	return nil
}
