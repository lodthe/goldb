package db

import (
	"context"
	"errors"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lodthe/goldb/db/dbclient"
	"github.com/stretchr/testify/assert"
)

func getBatch(start, batchSize int) []dbclient.Triplet {
	triplet := dbclient.Triplet{Lseq: "0000042"}
	resultTriplets := []dbclient.Triplet{}
	for id := start; id < start+batchSize; id++ {
		triplet.Key = "it" + strconv.Itoa(id)
		triplet.Value = strconv.Itoa(id)
		resultTriplets = append(resultTriplets, triplet)
	}
	return resultTriplets
}

func TestOneElmentIterator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := dbclient.NewMockClient(ctrl)

	conn, err := Open(WithClient(mock))
	assert.NoError(t, err, "failed to open connection")

	ctx := context.Background()
	iterator, err := newIterator(conn, ctx)
	assert.NoError(t, err, "failed to create iterator")

	triplet := dbclient.Triplet{Key: "it1", Value: "1", Lseq: "0000042"}
	resultTriplets := []dbclient.Triplet{triplet}
	mock.EXPECT().Seek(ctx, iterator.lseq, iterator.key, iterator.limit).Return(resultTriplets, nil)

	for iterator.HasNext() {
		item, err := iterator.GetNext()
		assert.NoError(t, err, "failed to get first item")

		dbTriplet := tripletFromInternal(triplet)
		assert.Equal(t, dbTriplet, item, "wrong frist item")

		mock.EXPECT().Seek(ctx, iterator.lseq, iterator.key, iterator.limit).Return([]dbclient.Triplet{}, nil)
	}
}

func TestExpectedErrorIterator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := dbclient.NewMockClient(ctrl)

	conn, err := Open(WithClient(mock))
	assert.NoError(t, err, "failed to open connection")

	ctx := context.Background()
	iterator, err := newIterator(conn, ctx)
	assert.NoError(t, err, "failed to create iterator")

	expectedError := errors.New("internal error")
	mock.EXPECT().Seek(ctx, iterator.lseq, iterator.key, iterator.limit).Return([]dbclient.Triplet{}, expectedError)

	_, err = iterator.GetNext()
	assert.ErrorIs(t, err, ErrIterationFinished, "error expected")
}
