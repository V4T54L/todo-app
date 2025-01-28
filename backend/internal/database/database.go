package database

// DB is an interface that abstracts the database operations.
// This allows you to mock the db connection in tests.
type DB interface {
	Ping() error
	Close() error
}
