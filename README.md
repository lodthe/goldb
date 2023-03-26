# GoLDB

Goldb is a ORM that simplifies interaction with [LSeqDB](https://github.com/ds-project-lseqdb/ds-project-public/tree/main).

It provides a robust interface for users and hides complicated logic of interaction with a multicluster LSM database.

## Examples

```go
import (
	"github.com/lodthe/goldb/db"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	conn, err := db.Open(
		db.WithLogger(logger),
		db.WithServerAddress("localhost:13337"),
	)
	if err != nil {
		logger.Error(err.Error())
	}

	defer conn.Close()
}
```

## Development

Use make to rebuild the protobuf configuration:

```bash
make build-proto
```
