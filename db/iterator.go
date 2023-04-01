package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/lodthe/goldb/db/dbclient"
	"go.uber.org/zap"
)

// DefaultBatchSize limits the maximum number of fetched items at a time.
// Such a limit is required to make network communication less loaded
// and escape gRPC message size limits.
const DefaultBatchSize = 1000

type Iterator struct {
	conn *Connection

	ctx context.Context

	finished bool

	currentBatch  []dbclient.Triplet
	batchPosition int

	lseq  string
	key   *string
	limit *uint32
}

func newIterator(conn *Connection, ctx context.Context, options ...IterOption) (*Iterator, error) {

	it := &Iterator{
		conn:          conn,
		ctx:           ctx,
		finished:      false,
		batchPosition: 0,
		currentBatch:  nil,
		lseq:          zeroVersion.lseq,
	}

	for idx, f := range options {
		err := f(it)
		if err != nil {
			return nil, fmt.Errorf("option no. %d cannot be applied: %w", idx, err)
		}
	}

	if it.limit == nil {
		var limit uint32 = DefaultBatchSize
		it.limit = &limit
	}

	return it, nil
}

// HasNext checks if the iterator contains still not read elements.
func (i *Iterator) HasNext() bool {
	if i.batchPosition != len(i.currentBatch) {
		return true
	}

	if i.finished {
		return false
	}

	err := i.loadNextBatch()
	if err != nil {
		i.conn.logger.Error("failed to load next batch", zap.Error(err))
		i.finished = true
	}

	return !i.finished
}

// GetNext returns the next triplet if iteration is not finished.
func (i *Iterator) GetNext() (Triplet, error) {
	if !i.HasNext() {
		return Triplet{}, ErrIterationFinished
	}

	i.batchPosition += 1

	return tripletFromInternal(i.currentBatch[i.batchPosition-1]), nil
}

func (i *Iterator) loadNextBatch() error {
	if i.finished {
		return nil
	}

	if i.batchPosition != len(i.currentBatch) {
		return nil
	}

	triplets, err := i.conn.client.Seek(i.ctx, i.lseq, i.key, i.limit)
	if err != nil {
		return err
	}

	i.conn.logger.Debug("SEEK succeed", zap.Int("count", len(triplets)))

	// Remove already seen records.
	if i.lseq != zeroVersion.lseq {
		for len(triplets) > 0 && triplets[0].Lseq <= i.lseq {
			triplets = triplets[1:]
		}
	}

	i.currentBatch = triplets
	i.batchPosition = 0

	// If records exist, forward the lseq iterator.
	if len(triplets) == 0 {
		i.finished = true
	} else {
		i.lseq = triplets[len(triplets)-1].Lseq
	}

	return nil
}

// IterOption configures an iterator.
// At the moment you can use only already written options.
type IterOption func(it *Iterator) error

// When IterKeyEquals is provided, all yielded triplets have the given key.
//
// By default triplets are not filtered by key.
func IterKeyEquals(key string) IterOption {
	return func(it *Iterator) error {
		if it.key != nil {
			return errors.New("key option cannot be specified twice")
		}

		it.key = &key

		return nil
	}
}

// When IterFromVersion is provided, all yielded triplets have version greater
// than the given one.
//
// By default there is no lower bound on allowed version.
func IterFromVersion(version Version) IterOption {
	return func(it *Iterator) error {
		if version.lseq > it.lseq {
			it.lseq = version.lseq
		}

		return nil
	}
}

// When IterSetLimit is provided, the given value limits the number
// of triplets fetched at a time from a server.
//
// By default the limit is set to DefaultBatchSize.
func IterSetLimit(limit *uint32) IterOption {
	return func(it *Iterator) error {
		if limit != nil && *limit == 0 {
			return errors.New("limit cannot be zero; nil must be provided to disable limiting")
		}

		if it.limit != nil {
			return errors.New("limit option cannot be specified twice")
		}

		it.limit = limit

		return nil
	}
}
