package db

import (
	"github.com/lodthe/goldb/db/dbclient"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Option = func(cfg *Config)

// WithLogger takes a logger instance will be used to log database operations.
// You can change the encoding and the log level to customize output.
//
// Default logger prints nothing.
func WithLogger(logger *zap.Logger) Option {
	return func(cfg *Config) {
		cfg.logger = logger
	}
}

// WithDebugLogger configures logger to be verbose and human-friendly
// to simplify the debug.
func WithDebugLogger() Option {
	return func(cfg *Config) {
		logger, _ := zap.NewDevelopment()
		cfg.logger = logger
	}
}

// WithClient takes an implementation of client that will be used to
// interact with a database server.
//
// Default client is missed that means a new gRPC client will be created.
// If this option is provided, options which control gRPC connection are ignored.
func WithClient(client dbclient.Client) Option {
	return func(cfg *Config) {
		cfg.client = client
	}
}

// WithConnection takes a gRPC connection that will be used to interact
// with a database server.
//
// If this option is missed, a new connection will be created from the
// given server address.
// If this option is provided, WithServerAddress is ignored.
func WithConnection(conn *grpc.ClientConn) Option {
	return func(cfg *Config) {
		cfg.grpcConn = conn
	}
}

// WithServerAddress takes address of the server gRPC API.
//
// The address is ignored if either a client or grpc connection are set by other options.
func WithServerAddress(address string) Option {
	return func(cfg *Config) {
		cfg.serverAddress = address
	}
}
