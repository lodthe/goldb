package dbclient

import "context"

type Client interface {
	Put(ctx context.Context, key, value string) (string, error)
	GetValue(ctx context.Context, key string) (string, string, error)

	CloseConnection() error
}
