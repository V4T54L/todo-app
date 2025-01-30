package services

import (
	"context"
	"todo_app_backend/internal/app/models"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, name string, email string, password string) (*models.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]models.User, error)
	GetUserByID(ctx context.Context, id int) (*models.User, error)
}

type TodoServiceInterface interface {
	CreateTodo(ctx context.Context, userId int, title string, content string) (*models.Todo, error)
	DeleteTodo(ctx context.Context, userId int, id int) error
	GetAllTodos(ctx context.Context, userId int) ([]models.Todo, error)
	GetTodoByID(ctx context.Context, userId int, id int) (*models.Todo, error)
	UpdateTodo(ctx context.Context, userId int, id int, title string, content string, status string) error
}
