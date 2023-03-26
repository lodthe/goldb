package dbclient

import (
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
