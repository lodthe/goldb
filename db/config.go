package db

import (
	"context"

	"github.com/lodthe/goldb/db/dbclient"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Config struct {
	ctx context.Context

	logger *zap.Logger

	client        dbclient.Client
	serverAddress string
	grpcConn      *grpc.ClientConn
}

var defaultConfig = Config{
	ctx:    context.Background(),
	logger: zap.NewNop(),
}
