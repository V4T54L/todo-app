package services

import (
	"context"
	"errors"
	"testing"

	"todo_app_backend/internal/app/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mockTodoRepo is a mock implementation of the TodoRepoInterface
type mockTodoRepo struct {
	mock.Mock
}

func (m *mockTodoRepo) CreateTodo(ctx context.Context, userId int, todo *models.Todo) error {
	args := m.Called(ctx, userId, todo)
	return args.Error(0)
}

func (m *mockTodoRepo) GetTodoByID(ctx context.Context, userId, id int) (*models.Todo, error) {
	args := m.Called(ctx, userId, id)
	return args.Get(0).(*models.Todo), args.Error(1)
}

func (m *mockTodoRepo) GetAllTodos(ctx context.Context, userId int) ([]models.Todo, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).([]models.Todo), args.Error(1)
}

func (m *mockTodoRepo) UpdateTodo(ctx context.Context, userId, id int, todo *models.Todo) error {
	args := m.Called(ctx, userId, id, todo)
	return args.Error(0)
}

func (m *mockTodoRepo) DeleteTodo(ctx context.Context, userId, id int) error {
	args := m.Called(ctx, userId, id)
	return args.Error(0)
}

func TestTodoService_CreateTodo(t *testing.T) {
	mockRepo := new(mockTodoRepo)
	service := NewTodoService(mockRepo)
	ctx := context.Background()

	// Test for success case
	todo := &models.Todo{Title: "Test Todo", Content: "Todo Content", Status: "pending"}
	mockRepo.On("CreateTodo", ctx, 1, todo).Return(nil)

	newTodo, err := service.CreateTodo(ctx, 1, "Test Todo", "Todo Content")
	assert.NoError(t, err)
	assert.Equal(t, todo.Title, newTodo.Title)

	// Test for error case due to empty title
	_, err = service.CreateTodo(ctx, 1, "", "Todo Content")
	assert.EqualError(t, err, "title is required")
}

func TestTodoService_GetTodoByID(t *testing.T) {
	mockRepo := new(mockTodoRepo)
	service := NewTodoService(mockRepo)
	ctx := context.Background()

	// Test for success case
	todo := &models.Todo{ID: 1, Title: "Test Todo", Content: "Todo Content"}
	mockRepo.On("GetTodoByID", ctx, 1, 1).Return(todo, nil)

	result, err := service.GetTodoByID(ctx, 1, 1)
	assert.NoError(t, err)
	assert.Equal(t, todo, result)

	// // Test for error case
	// mockRepo.On("GetTodoByID", ctx, 1, 1).Return(nil, errors.New("todo not found"))
	// result, err = service.GetTodoByID(ctx, 1, 1)
	// assert.Error(t, err)
	// assert.Nil(t, result)
}

func TestTodoService_GetTodoByIDFailure(t *testing.T) {
	mockRepo := new(mockTodoRepo)
	service := NewTodoService(mockRepo)
	ctx := context.Background()

	// Test for error case
	mockRepo.On("GetTodoByID", ctx, 1, 2).Return((*models.Todo)(nil), errors.New("todo not found"))
	result, err := service.GetTodoByID(ctx, 1, 2)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestTodoService_GetAllTodos(t *testing.T) {
	mockRepo := new(mockTodoRepo)
	service := NewTodoService(mockRepo)
	ctx := context.Background()

	// Test for success case
	todos := []models.Todo{
		{ID: 1, Title: "Test Todo 1"},
		{ID: 2, Title: "Test Todo 2"},
	}
	mockRepo.On("GetAllTodos", ctx, 1).Return(todos, nil)

	result, err := service.GetAllTodos(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, todos, result)
}

func TestTodoService_UpdateTodo(t *testing.T) {
	mockRepo := new(mockTodoRepo)
	service := NewTodoService(mockRepo)
	ctx := context.Background()

	// Test for success case
	todo := &models.Todo{Title: "Updated Todo", Content: "Updated Content", Status: "pending"}
	mockRepo.On("UpdateTodo", ctx, 1, 1, todo).Return(nil)

	err := service.UpdateTodo(ctx, 1, 1, "Updated Todo", "Updated Content", "pending")
	assert.NoError(t, err)

	// Test for error case due to empty title
	err = service.UpdateTodo(ctx, 1, 1, "", "Updated Content", "pending")
	assert.EqualError(t, err, "title is required")
}

func TestTodoService_DeleteTodo(t *testing.T) {
	mockRepo := new(mockTodoRepo)
	service := NewTodoService(mockRepo)
	ctx := context.Background()

	// Test for success case
	mockRepo.On("DeleteTodo", ctx, 1, 1).Return(nil)

	err := service.DeleteTodo(ctx, 1, 1)
	assert.NoError(t, err)

	// Test for error case
	mockRepo.On("DeleteTodo", ctx, 1, 2).Return(errors.New("todo not found"))
	err = service.DeleteTodo(ctx, 1, 2)
	assert.Error(t, err)
}
