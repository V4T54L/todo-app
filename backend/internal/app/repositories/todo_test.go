// todo_repo_test.go
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

func TestTodoRepository_CreateTodo(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := NewTodoRepository(mockDB)

	todo := &models.Todo{
		Title:   "Test Todo",
		Content: "This is a test",
		Status:  "Pending",
	}

	mock.ExpectQuery(`INSERT INTO todos .*`).
		WithArgs(1, todo.Title, todo.Content, todo.Status).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
			AddRow(1, time.Now(), time.Now()))

	err = repo.CreateTodo(context.Background(), 1, todo)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, todo.ID)

	// Test error scenario
	mock.ExpectQuery(`INSERT INTO todos .*`).WillReturnError(errors.New("db error"))
	err = repo.CreateTodo(context.Background(), 1, todo)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create todo")
}

func TestTodoRepository_GetTodoByID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := NewTodoRepository(mockDB)

	todo := &models.Todo{
		ID:      1,
		Title:   "Test Todo",
		Content: "This is a test",
		Status:  "Pending",
	}

	mock.ExpectQuery(`SELECT .* FROM todos WHERE id = \$1 AND user_id = \$2`).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content", "status", "created_at", "updated_at"}).
			AddRow(todo.ID, todo.Title, todo.Content, todo.Status, time.Now(), time.Now()))

	result, err := repo.GetTodoByID(context.Background(), 1, 1)
	assert.NoError(t, err)
	assert.Equal(t, todo.ID, result.ID)

	// Test not found scenario
	mock.ExpectQuery(`SELECT .* FROM todos WHERE id = \$1 AND user_id = \$2`).
		WithArgs(1, 1).
		WillReturnError(sql.ErrNoRows)
	_, err = repo.GetTodoByID(context.Background(), 1, 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "todo not found or does not belong to user")
}

func TestTodoRepository_GetAllTodos(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := NewTodoRepository(mockDB)

	rows := sqlmock.NewRows([]string{"id", "title", "content", "status", "created_at", "updated_at"}).
		AddRow(1, "Todo 1", "Content 1", "Pending", time.Now(), time.Now()).
		AddRow(2, "Todo 2", "Content 2", "Completed", time.Now(), time.Now())

	mock.ExpectQuery(`SELECT .* FROM todos WHERE user_id = \$1`).
		WithArgs(1).
		WillReturnRows(rows)

	todos, err := repo.GetAllTodos(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(todos))

	// Check if the returned todos match the expected values
	assert.Equal(t, "Todo 1", todos[0].Title)
	assert.Equal(t, "Content 1", todos[0].Content)
}

func TestTodoRepository_UpdateTodo(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := NewTodoRepository(mockDB)

	todo := &models.Todo{
		Title:   "Updated Todo",
		Content: "Updated Content",
		Status:  "Pending",
	}

	mock.ExpectExec(`UPDATE todos SET title = \$1, content = \$2, status = \$3 WHERE id = \$4 AND user_id = \$5`).
		WithArgs(todo.Title, todo.Content, todo.Status, 1, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.UpdateTodo(context.Background(), 1, 1, todo)
	assert.NoError(t, err)

	// Test not found scenario
	mock.ExpectExec(`UPDATE todos SET title = \$1, content = \$2, status = \$3 WHERE id = \$4 AND user_id = \$5`).
		WithArgs(todo.Title, todo.Content, todo.Status, 1, 1).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err = repo.UpdateTodo(context.Background(), 1, 1, todo)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "todo not found or does not belong to user")
}

func TestTodoRepository_DeleteTodo(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := NewTodoRepository(mockDB)

	mock.ExpectExec(`DELETE FROM todos WHERE id = \$1 AND user_id = \$2`).
		WithArgs(1, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.DeleteTodo(context.Background(), 1, 1)
	assert.NoError(t, err)

	// Test not found scenario
	mock.ExpectExec(`DELETE FROM todos WHERE id = \$1 AND user_id = \$2`).
		WithArgs(2, 1).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err = repo.DeleteTodo(context.Background(), 1, 2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "todo not found or does not belong to user")
}

func TestTodoRepository_GetAllTodos_Error(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := NewTodoRepository(mockDB)

	mock.ExpectQuery(`SELECT .* FROM todos WHERE user_id = \$1`).
		WithArgs(1).
		WillReturnError(errors.New("db error"))

	todos, err := repo.GetAllTodos(context.Background(), 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get all todos")
	assert.Nil(t, todos)
}
