// user_repo_test.go
package repositories

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"todo_app_backend/internal/app/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := NewUserRepository(mockDB)

	user := &models.User{
		Name:     "Test User",
		Email:    "testuser@example.com",
		Password: "securepassword",
	}

	mock.ExpectQuery(`INSERT INTO users .*`).
		WithArgs(user.Name, user.Email, user.Password).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	err = repo.CreateUser(context.Background(), user)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, user.ID)

	// Test error scenario
	mock.ExpectQuery(`INSERT INTO users .*`).WillReturnError(errors.New("db error"))
	err = repo.CreateUser(context.Background(), user)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create user")
}

func TestUserRepository_GetUserByID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := NewUserRepository(mockDB)

	userID := 1
	expectedUser := &models.User{
		ID:        userID,
		Name:      "Test User",
		Email:     "testuser@example.com",
		Password:  "securepassword",
		CreatedAt: "time.Now()",
	}

	mock.ExpectQuery(`SELECT id, name, email, password, created_at FROM users WHERE id = \$1`).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at"}).
			AddRow(expectedUser.ID, expectedUser.Name, expectedUser.Email, expectedUser.Password, expectedUser.CreatedAt))

	result, err := repo.GetUserByID(context.Background(), userID)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, result.ID)
	assert.Equal(t, expectedUser.Name, result.Name)

	// Test not found scenario
	mock.ExpectQuery(`SELECT id, name, email, password, created_at FROM users WHERE id = \$1`).
		WithArgs(userID).
		WillReturnError(sql.ErrNoRows)

	result, err = repo.GetUserByID(context.Background(), userID)
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.EqualError(t, err, "user not found")
}

func TestUserRepository_GetAllUsers(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := NewUserRepository(mockDB)

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at"}).
		AddRow(1, "User 1", "user1@example.com", "password1", time.Now()).
		AddRow(2, "User 2", "user2@example.com", "password2", time.Now())

	mock.ExpectQuery(`SELECT id, name, email, password, created_at FROM users`).
		WillReturnRows(rows)

	users, err := repo.GetAllUsers(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 2, len(users))
	assert.Equal(t, "User 1", users[0].Name)
	assert.Equal(t, "User 2", users[1].Name)

	// Test error scenario
	mock.ExpectQuery(`SELECT id, name, email, password, created_at FROM users`).
		WillReturnError(errors.New("db error"))

	users, err = repo.GetAllUsers(context.Background())
	assert.Error(t, err)
	assert.Nil(t, users)
	assert.Contains(t, err.Error(), "failed to get all users")
}

func TestUserRepository_DeleteUser(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := NewUserRepository(mockDB)

	userID := 1

	mock.ExpectExec(`DELETE FROM users WHERE id = \$1`).
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(0, 1)) // 1 row deleted

	err = repo.DeleteUser(context.Background(), userID)
	assert.NoError(t, err)

	// Test not found scenario
	mock.ExpectExec(`DELETE FROM users WHERE id = \$1`).
		WithArgs(userID + 1).
		WillReturnResult(sqlmock.NewResult(0, 0)) // No rows deleted

	err = repo.DeleteUser(context.Background(), userID+1)
	assert.Error(t, err)
	assert.EqualError(t, err, "user not found")

	// Test execution error scenario
	mock.ExpectExec(`DELETE FROM users WHERE id = \$1`).
		WithArgs(userID).
		WillReturnError(errors.New("db error"))

	err = repo.DeleteUser(context.Background(), userID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to delete user")
}
