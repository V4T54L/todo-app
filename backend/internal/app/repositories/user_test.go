package repositories

import (
	"database/sql"
	"testing"

	"todo_app_backend/internal/app/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)
	user := &models.User{Name: "John Doe", Email: "john@example.com", Password: "password123"}

	mock.ExpectQuery(`INSERT INTO users \(name, email, password\) VALUES \(\$1, \$2, \$3\) RETURNING id`).
		WithArgs(user.Name, user.Email, user.Password).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	err = repo.CreateUser(user)
	require.NoError(t, err)
	assert.Equal(t, 1, user.ID)
}

func TestGetUserByID_UserExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)
	mock.ExpectQuery(`SELECT id, name, email, password, created_at FROM users WHERE id = \$1`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at"}).
			AddRow(1, "John Doe", "john@example.com", "password123", "2023-01-01 10:00:00"))

	user, err := repo.GetUserByID(1)
	require.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "John Doe", user.Name)
}

func TestGetUserByID_UserNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)
	mock.ExpectQuery(`SELECT id, name, email, password, created_at FROM users WHERE id = \$1`).
		WithArgs(999).
		WillReturnError(sql.ErrNoRows)

	user, err := repo.GetUserByID(999)
	require.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "user not found", err.Error())
}

func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)
	mock.ExpectQuery(`SELECT id, name, email, password, created_at FROM users`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at"}).
			AddRow(1, "John Doe", "john@example.com", "password123", "2023-01-01 10:00:00").
			AddRow(2, "Jane Doe", "jane@example.com", "password123", "2023-01-02 10:00:00"))

	users, err := repo.GetAllUsers()
	require.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, "John Doe", users[0].Name)
	assert.Equal(t, "Jane Doe", users[1].Name)
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)
	mock.ExpectExec(`DELETE FROM users WHERE id = \$1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.DeleteUser(1)
	require.NoError(t, err)
}
