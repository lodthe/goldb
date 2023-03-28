package dbclient

import (
	"context"

	"github.com/lodthe/goldb/proto/lseqdbpb"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	conn *grpc.ClientConn
	cli  lseqdbpb.LSeqDatabaseClient
}

func NewGRPCClient(conn *grpc.ClientConn) *GRPCClient {
	return &GRPCClient{
		conn: conn,
		cli:  lseqdbpb.NewLSeqDatabaseClient(conn),
	}
}

func (c *GRPCClient) CloseConnection() error {
	return c.conn.Close()
}

func (c *GRPCClient) Put(ctx context.Context, key, value string) (string, error) {
	r, err := c.cli.Put(ctx, &lseqdbpb.PutRequest{
		Key:   key,
		Value: value,
	})
	if err != nil {
		return "", err
	}

	return r.Lseq, nil
}

func (c *GRPCClient) GetValue(ctx context.Context, key string) (string, string, error) {
	r, err := c.cli.GetValue(ctx, &lseqdbpb.ReplicaKey{
		Key: key,
	})
	if err != nil {
		return "", "", err
	}

	return r.Value, r.Lseq, nil
}
