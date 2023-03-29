package dbclient

import (
	"context"

	"github.com/lodthe/goldb/proto/lseqdbpb"
	"go.uber.org/zap"
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

func (c *GRPCClient) Seek(ctx context.Context, lseq string, key *string, limit *uint32) ([]Triplet, error) {
	r, err := c.cli.SeekGet(ctx, &lseqdbpb.SeekGetRequest{
		Key:   key,
		Lseq:  lseq,
		Limit: limit,
	})

	if err != nil {
		l, _ := zap.NewDevelopment()
		l.Sugar().Info(err.Error())
	}

	if err != nil {
		return nil, err
	}

	triplets := make([]Triplet, 0, len(r.Items))
	for _, i := range r.Items {
		if i == nil {
			continue
		}

		triplets = append(triplets, Triplet{
			Key:   i.Key,
			Value: i.Value,
			Lseq:  i.Lseq,
		})
	}

	return triplets, nil
}
