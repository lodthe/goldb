# GoLDB

Goldb is a ORM that simplifies interaction with [LSeqDB](https://github.com/ds-project-lseqdb/ds-project-public/tree/main).

It provides a robust interface for users and hides complicated logic of interaction with a multicluster LSM database.

## Examples

See the [examples](./examples/) directory for sources.

### Connection initialization

```go
ctx := context.Background()
logger, _ := zap.NewDevelopment()

// Establish a connection with the server.
conn, err := db.Open(
    db.WithLogger(logger),
    // Provide server address here.
    db.WithServerAddress("bloom.lodthe.me:8888"),
)
if err != nil {
    log.Fatal("failed to connect:", err)
}

defer conn.Close()
```

### Data model

Database record is represented as a triplet:

```go
type Triplet struct {
    Key     string
    Value   string
    Version Version
}
```

`Version` is a unique value that represents internal partial-ordered version (Lamport's sequence number).

### Put / GetLatest

```go
key := "Alice"
value := "Alice's shopping cart"

// Create a new record.
triplet, err := conn.Put(ctx, key, value)
if err != nil {
    log.Fatal("put failed:", err)
}

// Get the latest value for "Alice" key.
triplet, err = conn.GetLatest(ctx, key)
if err != nil {
    log.Fatal("get latest failed:", err)
}

log.Printf("got values: %s -> %s (%s)", triplet.Key, triplet.Value, triplet.Version)
```

## Iterator

It is possible to iterate over data:

```go
options := []db.IterOption{
    // Get only triplets with "Alice" key.
    // If no options provided, all triplets will be returned.
    db.IterKeyEquals("Alice"),
}

iterator, err := conn.GetIterator(context.Background(), options...)
if err != nil {
    logger.Fatal(err.Error())
}

for iterator.HasNext() {
    item, err := iterator.GetNext()
    // Handle error.

    log.Printf("[%s] %s -> %s", item.Version, item.Key, item.Value)
}
```

Iterator takes a set of options that can set constraints on the returned data.

## Development

Use make to rebuild the protobuf configuration:

```bash
make build-proto
```
