package v1_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	v1 "todo_app_backend/api/v1"
	"todo_app_backend/internal/app/models"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTodoService is a mock implementation of TodoServiceInterface
type MockTodoService struct {
	mock.Mock
}

func (m *MockTodoService) CreateTodo(ctx context.Context, userId int, title, content string) (*models.Todo, error) {
	args := m.Called(ctx, userId, title, content)
	return args.Get(0).(*models.Todo), args.Error(1)
}

func (m *MockTodoService) GetAllTodos(ctx context.Context, userId int) ([]models.Todo, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).([]models.Todo), args.Error(1)
}

func (m *MockTodoService) GetTodoByID(ctx context.Context, userId, id int) (*models.Todo, error) {
	args := m.Called(ctx, userId, id)
	return args.Get(0).(*models.Todo), args.Error(1)
}

func (m *MockTodoService) UpdateTodo(ctx context.Context, userId, id int, title, content, status string) error {
	args := m.Called(ctx, userId, id, title, content, status)
	return args.Error(0)
}

func (m *MockTodoService) DeleteTodo(ctx context.Context, userId, id int) error {
	args := m.Called(ctx, userId, id)
	return args.Error(0)
}

func TestCreateTodo(t *testing.T) {
	mockService := new(MockTodoService)
	handler := v1.NewTodoHandler(mockService)

	todo := models.Todo{Title: "Test Todo", Content: "Content of test todo"}
	mockService.On("CreateTodo", mock.Anything, 1, todo.Title, todo.Content).Return(&todo, nil)

	body, _ := json.Marshal(todo)
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(body))
	req = req.WithContext(context.WithValue(req.Context(), "userID", 1))
	resp := httptest.NewRecorder()

	handler.CreateTodo(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var createdTodo models.Todo
	json.NewDecoder(resp.Body).Decode(&createdTodo)
	assert.Equal(t, todo.Title, createdTodo.Title)
	assert.Equal(t, todo.Content, createdTodo.Content)
}

func TestCreateTodo_BadRequest(t *testing.T) {
	mockService := new(MockTodoService)
	handler := v1.NewTodoHandler(mockService)

	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer([]byte("invalid json")))
	resp := httptest.NewRecorder()

	handler.CreateTodo(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestGetAllTodos_Success(t *testing.T) {
	mockService := new(MockTodoService)
	handler := v1.NewTodoHandler(mockService)

	todos := []models.Todo{
		{ID: 1, Title: "Todo 1", Content: "Content 1"},
		{ID: 2, Title: "Todo 2", Content: "Content 2"},
	}
	mockService.On("GetAllTodos", mock.Anything, 1).Return(todos, nil)

	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	req = req.WithContext(context.WithValue(req.Context(), "userID", 1))
	resp := httptest.NewRecorder()

	handler.GetAllTodos(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var receivedTodos []models.Todo
	json.NewDecoder(resp.Body).Decode(&receivedTodos)
	assert.Equal(t, len(todos), len(receivedTodos))
}

func TestGetTodoByID_Success(t *testing.T) {
	mockService := new(MockTodoService)
	handler := v1.NewTodoHandler(mockService)

	todo := models.Todo{ID: 1, Title: "Todo 1", Content: "Content 1"}
	mockService.On("GetTodoByID", mock.Anything, 1, 1).Return(&todo, nil)

	req := httptest.NewRequest(http.MethodGet, "/todos/1", nil)
	req = req.WithContext(context.WithValue(req.Context(), "userID", 1))
	resp := httptest.NewRecorder()
	router := chi.NewRouter()
	router.Get("/todos/{id}", handler.GetTodoByID)

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var receivedTodo models.Todo
	json.NewDecoder(resp.Body).Decode(&receivedTodo)
	assert.Equal(t, todo.Title, receivedTodo.Title)
}

func TestGetTodoByID_NotFound(t *testing.T) {
	mockService := new(MockTodoService)
	handler := v1.NewTodoHandler(mockService)

	mockService.On("GetTodoByID", mock.Anything, 1, 999).Return(&models.Todo{}, errors.New("todo not founs"))

	req := httptest.NewRequest(http.MethodGet, "/todos/999", nil)
	req = req.WithContext(context.WithValue(req.Context(), "userID", 1))
	resp := httptest.NewRecorder()
	router := chi.NewRouter()
	router.Get("/todos/{id}", handler.GetTodoByID)

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestUpdateTodo_Success(t *testing.T) {
	mockService := new(MockTodoService)
	handler := v1.NewTodoHandler(mockService)

	todo := models.Todo{Title: "Updated Todo", Content: "Updated Content"}
	mockService.On("UpdateTodo", mock.Anything, 1, 1, todo.Title, todo.Content, "").Return(nil)

	body, _ := json.Marshal(todo)
	req := httptest.NewRequest(http.MethodPut, "/todos/{id}", bytes.NewBuffer(body))
	req = req.WithContext(context.WithValue(req.Context(), "userID", 1))
	resp := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.UpdateTodo(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
}

func TestUpdateTodo_BadRequest(t *testing.T) {
	mockService := new(MockTodoService)
	handler := v1.NewTodoHandler(mockService)

	req := httptest.NewRequest(http.MethodPut, "/todos/{id}", bytes.NewBuffer([]byte("{'invalid':'json'")))
	resp := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.UpdateTodo(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "invalid json", strings.Trim(string(body), "\n"))
}

func TestDeleteTodo_Success(t *testing.T) {
	mockService := new(MockTodoService)
	handler := v1.NewTodoHandler(mockService)

	mockService.On("DeleteTodo", mock.Anything, 1, 1).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/todos/{id}", nil)
	req = req.WithContext(context.WithValue(req.Context(), "userID", 1))
	resp := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.DeleteTodo(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
}

func TestDeleteTodo_NotFound(t *testing.T) {
	mockService := new(MockTodoService)
	handler := v1.NewTodoHandler(mockService)

	mockService.On("DeleteTodo", mock.Anything, 1, 999).Return(errors.New("todo not founs"))

	req := httptest.NewRequest(http.MethodDelete, "/todos/{id}", nil)
	req = req.WithContext(context.WithValue(req.Context(), "userID", 1))
	resp := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "999")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.DeleteTodo(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}
