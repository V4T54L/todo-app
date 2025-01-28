// db/postgres_db.go
package database

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

// PostgreSQLDB implements the DB interface
type PostgreSQLDB struct {
	conn *sql.DB
}

// NewPostgreSQLDB initializes a new PostgreSQLDB instance
func NewPostgreSQLDB(dbUri string, maxIdleConns, maxOpenConns int) (*PostgreSQLDB, error) {
	dbConn, err := sql.Open("postgres", dbUri)
	if err != nil {
		return nil, err
	}

	dbConn.SetMaxOpenConns(maxOpenConns)
	dbConn.SetMaxIdleConns(maxIdleConns)

	if err = dbConn.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to PostgreSQL database")
	return &PostgreSQLDB{conn: dbConn}, nil
}

// GetConn returns the underlying sql.DB connection
func (p *PostgreSQLDB) GetConn() *sql.DB {
	return p.conn
}

// Ping pings the database connection
func (p *PostgreSQLDB) Ping() error {
	return p.conn.Ping()
}

// Close closes the database connection
func (p *PostgreSQLDB) Close() error {
	if p.conn == nil {
		return errors.New("database connection is nil")
	}
	return p.conn.Close()
}

// Ensure PostgreSQLDB implements the DB interface
var _ DB = (*PostgreSQLDB)(nil)
