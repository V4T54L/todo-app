// todo_repo.go
package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"todo_app_backend/internal/app/models"
)

type TodoRepository struct {
	DB *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

// CreateTodo inserts a new todo into the database
func (r *TodoRepository) CreateTodo(ctx context.Context, userId int, todo *models.Todo) error {
	query := `INSERT INTO todos (user_id, title, content, status) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at;`
	err := r.DB.QueryRowContext(ctx, query, userId, todo.Title, todo.Content, todo.Status).
		Scan(&todo.ID, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create todo: %w", err)
	}
	return nil
}

// GetTodoByID retrieves a todo by ID
func (r *TodoRepository) GetTodoByID(ctx context.Context, userId, id int) (*models.Todo, error) {
	todo := &models.Todo{}
	query := `SELECT id, title, content, status, created_at, updated_at FROM todos WHERE id = $1 AND user_id = $2`
	err := r.DB.QueryRowContext(ctx, query, id, userId).
		Scan(&todo.ID, &todo.Title, &todo.Content, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("todo not found or does not belong to user")
	} else if err != nil {
		return nil, fmt.Errorf("failed to get todo by ID: %w", err)
	}

	return todo, nil
}

// GetAllTodos retrieves all todos
func (r *TodoRepository) GetAllTodos(ctx context.Context, userId int) ([]models.Todo, error) {
	var todos []models.Todo
	query := `SELECT id, title, content, status, created_at, updated_at FROM todos WHERE user_id = $1`
	rows, err := r.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get all todos: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Content, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan todo: %w", err)
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return todos, nil
}

// UpdateTodo updates the todo with the provided ID
func (r *TodoRepository) UpdateTodo(ctx context.Context, userId, id int, todo *models.Todo) error {
	query := `UPDATE todos SET title = $1, content = $2, status = $3 WHERE id = $4 AND user_id = $5`
	result, err := r.DB.ExecContext(ctx, query, todo.Title, todo.Content, todo.Status, id, userId)
	if err != nil {
		return fmt.Errorf("failed to update todo: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected during update: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("todo not found or does not belong to user")
	}

	return nil
}

// DeleteTodo removes a todo by ID
func (r *TodoRepository) DeleteTodo(ctx context.Context, userId, id int) error {
	query := `DELETE FROM todos WHERE id = $1 AND user_id = $2`
	result, err := r.DB.ExecContext(ctx, query, id, userId)
	if err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected during delete: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("todo not found or does not belong to user")
	}

	return nil
}

var _ TodoRepoInterface = (*TodoRepository)(nil)
