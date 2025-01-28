// todo_repo.go
package repositories

import (
	"database/sql"
	"errors"
	"todo_app_backend/internal/app/models"
)

type TodoRepository struct {
	DB *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

// CreateTodo inserts a new todo into the database
func (r *TodoRepository) CreateTodo(todo *models.Todo) error {
	query := `INSERT INTO todos (title, content, done) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at;`
	err := r.DB.QueryRow(query, todo.Title, todo.Content, todo.Done).Scan(&todo.ID, &todo.CreatedAt, &todo.UpdatedAt)
	return err
}

// GetTodoByID retrieves a todo by ID
func (r *TodoRepository) GetTodoByID(id int) (*models.Todo, error) {
	todo := &models.Todo{}
	query := `SELECT id, title, content, done, created_at, updated_at FROM todos WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&todo.ID, &todo.Title, &todo.Content, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("todo not found")
	}
	return todo, err
}

// GetAllTodos retrieves all todos
func (r *TodoRepository) GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo
	query := `SELECT id, title, content, done, created_at, updated_at FROM todos`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Content, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

// UpdateTodo updates the todo with the provided ID
func (r *TodoRepository) UpdateTodo(id int, todo *models.Todo) error {
	query := `UPDATE todos SET title = $1, content = $2, done = $3 WHERE id = $4`
	_, err := r.DB.Exec(query, todo.Title, todo.Content, todo.Done, id)
	return err
}

// DeleteTodo removes a todo by ID
func (r *TodoRepository) DeleteTodo(id int) error {
	query := `DELETE FROM todos WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
