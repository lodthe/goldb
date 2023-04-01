package db

import (
	"context"
	"errors"
	"fmt"
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

func TestIteratorOptions(t *testing.T) {
	var limit uint32 = 42
	options := []struct {
		iterOption IterOption
		name       string
	}{
		{
			iterOption: IterKeyEquals("it1"),
			name:       "KeyOption",
		},
		{
			iterOption: IterFromVersion(NewVersion("0000042")),
			name:       "VersionOption",
		},
		{
			iterOption: IterSetLimit(&limit),
			name:       "LimitOption",
		},
	}

	for _, test := range options {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := dbclient.NewMockClient(ctrl)

			conn, err := Open(WithClient(mock))
			assert.NoError(t, err, "failed to open connection")

			ctx := context.Background()
			_, err = newIterator(conn, ctx, test.iterOption)
			assert.NoError(t, err, "failed to create iterator")
		})
	}
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

func TestOneAndHalfBatchIterator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := dbclient.NewMockClient(ctrl)

	conn, err := Open(WithClient(mock))
	assert.NoError(t, err, "failed to open connection")

	ctx := context.Background()
	iterator, err := newIterator(conn, ctx)
	assert.NoError(t, err, "failed to create iterator")

	resultTriplets := getBatch(0, DefaultBatchSize)
	mock.EXPECT().Seek(ctx, iterator.lseq, iterator.key, iterator.limit).Return(resultTriplets, nil)

	step := 0
	for iterator.HasNext() {
		item, err := iterator.GetNext()
		assert.NoError(t, err, fmt.Sprintf("failed to get item after %d step", step))

		dbTriplet := tripletFromInternal(resultTriplets[step])
		assert.Equal(t, dbTriplet, item, fmt.Sprintf("wrong item after %d step", step))

		step++

		if step == len(resultTriplets) {
			if len(resultTriplets) == DefaultBatchSize {
				secondBatch := getBatch(DefaultBatchSize, DefaultBatchSize/2)
				mock.EXPECT().Seek(ctx, iterator.lseq, iterator.key, iterator.limit).Return(secondBatch, nil)
				resultTriplets = append(resultTriplets, secondBatch...)
			} else {
				mock.EXPECT().Seek(ctx, iterator.lseq, iterator.key, iterator.limit).Return([]dbclient.Triplet{}, nil)
			}
		}
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
