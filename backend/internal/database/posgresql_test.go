// internal/infrastructure/db/postgresql_test.go
package database

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPostgreSQLDB(t *testing.T) {
	// Set up a mock database connection
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	require.NoError(t, err)
	defer db.Close()

	// Use sqlmock's ExpectPing to simulate the database pinging
	mock.ExpectPing().WillReturnError(nil)
	mock.ExpectClose()

	// Create a new PostgreSQLDB instance
	pgDB := &PostgreSQLDB{conn: db}

	// Test the successful ping of the database
	err = pgDB.Ping()
	assert.NoError(t, err)

	// Close the database connection
	err = pgDB.Close()
	assert.NoError(t, err)
}

func TestNewPostgreSQLDB_Error(t *testing.T) {
	// Test the error scenario when the connection cannot be opened
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	require.NoError(t, err)
	defer db.Close()

	// Simulate error on Ping
	mock.ExpectPing().WillReturnError(errors.New("database connection error"))
	mock.ExpectClose()

	pgDB := &PostgreSQLDB{conn: db}

	// Test the error when pinging the database
	err = pgDB.Ping()
	assert.Error(t, err)
	assert.Equal(t, "database connection error", err.Error())
}

func TestClose(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	require.NoError(t, err)
	defer db.Close()

	// Expect a call to Close
	mock.ExpectClose()

	pgDB := &PostgreSQLDB{conn: db}

	// Call Close and assert no error
	err = pgDB.Close()
	assert.NoError(t, err)

	// Ensure that all expectations were met
	require.NoError(t, mock.ExpectationsWereMet())

	// Attempt to close a nil connection
	pgDB.conn = nil
	err = pgDB.Close()
	assert.Error(t, err)
	assert.Equal(t, "database connection is nil", err.Error())
}

func TestGetConn(t *testing.T) {
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	pgDB := &PostgreSQLDB{conn: db}
	conn := pgDB.GetConn()

	assert.NotNil(t, conn)
}
