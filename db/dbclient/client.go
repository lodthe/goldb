package dbclient

type Client interface {
	CloseConnection() error
}
