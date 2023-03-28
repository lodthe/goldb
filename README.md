# GoLDB

Goldb is a ORM that simplifies interaction with [LSeqDB](https://github.com/ds-project-lseqdb/ds-project-public/tree/main).

It provides a robust interface for users and hides complicated logic of interaction with a multicluster LSM database.

## Examples

```go
package main

import (
	"context"
	"fmt"

	"github.com/lodthe/goldb/db"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	logger, _ := zap.NewDevelopment()

	conn, err := db.Open(
		db.WithLogger(logger),
		db.WithServerAddress("bloom.lodthe.me:8888"),
	)
	if err != nil {
		logger.Error(err.Error())
	}

	defer conn.Close()

	key := "user#10"
    value := "name:John"

	triplet, err := conn.Put(ctx, key, value)
	if err != nil {
		logger.Fatal("failed to put", zap.Error(err))
	}

	triplet, err = conn.GetLatest(ctx, key)
	if err != nil {
		logger.Fatal("failed to get latest", zap.Error(err))
	}

	logger.Sugar().Infof("got values: %s -> %s (%s)", triplet.Key, triplet.Value, triplet.Version)

    // Output:
    // got values: user#10 -> name:John (#000000002000000000000005)
}

```

## Development

Use make to rebuild the protobuf configuration:

```bash
make build-proto
```
