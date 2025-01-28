package repositories

import (
	"database/sql"
	"testing"
	"time"

	"todo_app_backend/internal/app/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewTodoRepository(db)

	todo := &models.Todo{
		Title:   "Test Todo",
		Content: "Test Content",
		Done:    false,
	}

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(1, time.Now(), time.Now())
	mock.ExpectQuery(`INSERT INTO todos`).WithArgs(todo.Title, todo.Content, todo.Done).
		WillReturnRows(rows)

	err = repo.CreateTodo(todo)
	require.NoError(t, err)
	assert.NotZero(t, todo.ID) // Verify that ID is set
}

func TestGetTodoByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewTodoRepository(db)

	rows := sqlmock.NewRows([]string{"id", "title", "content", "done", "created_at", "updated_at"}).
		AddRow(1, "Test Todo", "Test Content", false, time.Now(), time.Now())

	mock.ExpectQuery(`SELECT id, title, content, done, created_at, updated_at FROM todos WHERE id = \$1`).
		WithArgs(1).
		WillReturnRows(rows)

	todo, err := repo.GetTodoByID(1)
	require.NoError(t, err)
	assert.NotNil(t, todo)
	assert.Equal(t, 1, todo.ID)
}

func TestGetTodoByIDNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewTodoRepository(db)

	mock.ExpectQuery(`SELECT id, title, content, done, created_at, updated_at FROM todos WHERE id = \$1`).
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)

	todo, err := repo.GetTodoByID(1)
	require.Error(t, err)
	assert.Nil(t, todo)
	assert.Equal(t, "todo not found", err.Error())
}

func TestGetAllTodos(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewTodoRepository(db)

	rows := sqlmock.NewRows([]string{"id", "title", "content", "done", "created_at", "updated_at"}).
		AddRow(1, "Todo 1", "Content 1", false, time.Now(), time.Now()).
		AddRow(2, "Todo 2", "Content 2", true, time.Now(), time.Now())

	mock.ExpectQuery(`SELECT id, title, content, done, created_at, updated_at FROM todos`).
		WillReturnRows(rows)

	todos, err := repo.GetAllTodos()
	require.NoError(t, err)
	assert.Len(t, todos, 2)
}

func TestUpdateTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewTodoRepository(db)

	todo := &models.Todo{
		Title:   "Updated Todo",
		Content: "Updated Content",
		Done:    true,
	}

	mock.ExpectExec(`UPDATE todos SET title = \$1, content = \$2, done = \$3 WHERE id = \$4`).
		WithArgs(todo.Title, todo.Content, todo.Done, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateTodo(1, todo)
	require.NoError(t, err)
}

func TestDeleteTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewTodoRepository(db)

	mock.ExpectExec(`DELETE FROM todos WHERE id = \$1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.DeleteTodo(1)
	require.NoError(t, err)
}
