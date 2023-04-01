package db

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lodthe/goldb/db/dbclient"
	"github.com/stretchr/testify/assert"
)

func TestPut(t *testing.T) {
	testcases := []struct {
		name string

		key   string
		value string

		responseLseq string
		responseErr  error
	}{
		{
			name:         "OK test",
			key:          "Alice",
			value:        "apples,oranges",
			responseLseq: "0000042",
			responseErr:  nil,
		},
		{
			name:         "Internal error",
			key:          "Alice",
			value:        "apples,oranges",
			responseLseq: "",
			responseErr:  errors.New("internal error"),
		},
	}

	for _, test := range testcases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := dbclient.NewMockClient(ctrl)

			conn, err := Open(WithClient(mock))
			assert.NoError(t, err, "failed to open connection")

			ctx := context.Background()
			mock.EXPECT().Put(ctx, test.key, test.value).Return(test.responseLseq, test.responseErr)

			triplet, err := conn.Put(ctx, test.key, test.value)
			if test.responseErr != nil {
				assert.ErrorIs(t, err, test.responseErr, "error expected")
				return
			}

			assert.NoError(t, err, "failed to put kv")
			assert.Equal(t, test.key, triplet.Key, "invalid key")
			assert.Equal(t, test.value, triplet.Value, "invalid value")
			assert.Equal(t, test.responseLseq, triplet.Version.lseq, "invalid version")
		})
	}
}
