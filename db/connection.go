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

// Put saves a new pair and returns the version of the inserted data in case of success.
func (c *Connection) Put(ctx context.Context, key string, value string) (Triplet, error) {
	lseq, err := c.client.Put(ctx, key, value)
	if err != nil {
		c.logger.Warn("PUT failed", zap.String("key", key))

		return Triplet{}, fmt.Errorf("Put failed: %w", err)
	}

	c.logger.Debug("PUT succeed", zap.String("key", key), zap.String("version", lseq))

	return Triplet{
		Key:     key,
		Value:   value,
		Version: newVersion(lseq),
	}, nil
}

// GetLast finds a triplet by key and returns the latest known
// from the perspective of the replica to which a connection
// has been established.
//
// Triplets are sorted by version.
func (c *Connection) GetLatest(ctx context.Context, key string) (Triplet, error) {
	value, lseq, err := c.client.GetValue(ctx, key)
	if err != nil {
		c.logger.Warn("GET_LATEST failed", zap.String("key", key))

		return Triplet{}, fmt.Errorf("GetValue failed: %w", err)
	}

	c.logger.Debug("GET_LATEST succeed", zap.String("key", key), zap.String("version", lseq))

	return Triplet{
		Key:     key,
		Value:   value,
		Version: newVersion(lseq),
	}, nil
}

// GetIterator creates an iterator to iterate over values that satisfy the provided
// options.
//
// See the iterator.go file to get list of available filters:
// - IterKeyEquals
// - IterFromVersion
// - IterSetLimit
func (c *Connection) GetIterator(ctx context.Context, options ...IterOption) (*Iterator, error) {
	iterator, err := newIterator(c, ctx, options...)
	if err != nil {
		return nil, fmt.Errorf("iterator cannot be created: %w", err)
	}

	return iterator, nil
}
