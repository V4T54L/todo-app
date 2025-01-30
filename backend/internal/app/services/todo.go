package services

import (
	"context"
	"errors"
	"todo_app_backend/internal/app/models"
	"todo_app_backend/internal/app/repositories"
)

// TodoService defines methods related to todo operations.
type TodoService struct {
	TodoRepo repositories.TodoRepoInterface
}

// NewTodoService initializes a new TodoService.
func NewTodoService(todoRepo repositories.TodoRepoInterface) *TodoService {
	return &TodoService{TodoRepo: todoRepo}
}

// CreateTodo creates a new todo.
func (s *TodoService) CreateTodo(ctx context.Context, userId int, title, content string) (*models.Todo, error) {
	// Validate input
	if title == "" {
		return nil, errors.New("title is required")
	}

	// Create todo model
	todo := &models.Todo{
		Title:   title,
		Content: content,
		Status:  "pending",
	}

	// Store todo in DB
	err := s.TodoRepo.CreateTodo(ctx, userId, todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// GetTodoByID retrieves a todo by ID.
func (s *TodoService) GetTodoByID(ctx context.Context, userId, id int) (*models.Todo, error) {
	todo, err := s.TodoRepo.GetTodoByID(ctx, userId, id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// GetAllTodos retrieves all todos.
func (s *TodoService) GetAllTodos(ctx context.Context, userId int) ([]models.Todo, error) {
	return s.TodoRepo.GetAllTodos(ctx, userId)
}

// UpdateTodo updates a todo by ID.
func (s *TodoService) UpdateTodo(ctx context.Context, userId int, id int, title, content, status string) error {
	if title == "" {
		return errors.New("title is required")
	}

	todo := &models.Todo{
		Title:   title,
		Content: content,
		Status:  status,
	}

	return s.TodoRepo.UpdateTodo(ctx, userId, id, todo)
}

// DeleteTodo deletes a todo by ID.
func (s *TodoService) DeleteTodo(ctx context.Context, userId int, id int) error {
	return s.TodoRepo.DeleteTodo(ctx, userId, id)
}

var _ TodoServiceInterface = (*TodoService)(nil)
