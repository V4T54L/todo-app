package repositories

import (
	"context"
	"todo_app_backend/internal/app/models"
)

type TodoRepoInterface interface {
	CreateTodo(ctx context.Context, userId int, todo *models.Todo) error
	DeleteTodo(ctx context.Context, userId int, id int) error
	GetAllTodos(ctx context.Context, userId int) ([]models.Todo, error)
	GetTodoByID(ctx context.Context, userId int, id int) (*models.Todo, error)
	UpdateTodo(ctx context.Context, userId int, id int, todo *models.Todo) error
}

type UserRepoInterface interface {
	CreateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]models.User, error)
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}
