package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"todo_app_backend/internal/app/models"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock implementation of UserServiceInterface.
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(ctx context.Context, name, email, password string) (*models.User, error) {
	args := m.Called(ctx, name, email, password)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserService) DeleteUser(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	reqBody := models.User{Name: "John Doe", Email: "john@example.com", Password: "password123"}
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(jsonBody))
	rr := httptest.NewRecorder()

	mockService.On("CreateUser", mock.Anything, reqBody.Name, reqBody.Email, reqBody.Password).Return(&models.User{ID: 1, Name: reqBody.Name, Email: reqBody.Email}, nil)

	handler.CreateUser(rr, req)

	res := rr.Result()
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	var user models.User
	json.NewDecoder(res.Body).Decode(&user)
	assert.Equal(t, reqBody.Name, user.Name)
	assert.Equal(t, reqBody.Email, user.Email)
}

func TestCreateUser_InvalidJSON(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(`{"invalid": "json"`)))
	rr := httptest.NewRecorder()

	handler.CreateUser(rr, req)

	res := rr.Result()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestGetUserByID(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users/{id}", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	mockService.On("GetUserByID", mock.Anything, 1).Return(&models.User{ID: 1, Name: "John Doe", Email: "john@example.com"}, nil)

	handler.GetUserByID(rr, req)

	res := rr.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var user models.User
	json.NewDecoder(res.Body).Decode(&user)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "John Doe", user.Name)
}

func TestGetUserByID_InvalidID(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users/{id}", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "invalid")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.GetUserByID(rr, req)

	res := rr.Result()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestGetAllUsers(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rr := httptest.NewRecorder()

	mockService.On("GetAllUsers", mock.Anything).Return([]models.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
	}, nil)

	handler.GetAllUsers(rr, req)

	res := rr.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var users []models.User
	json.NewDecoder(res.Body).Decode(&users)
	assert.Len(t, users, 2)
}

func TestDeleteUser(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	rr := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	mockService.On("DeleteUser", mock.Anything, 1).Return(fmt.Errorf("user not found")).Once()

	handler.DeleteUser(rr, req)

	res := rr.Result()
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	assert.Equal(t, "user not found", strings.Trim(string(body), "\n"))
}

func TestDeleteUser_InvalidID(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/users/invalid", nil)
	rr := httptest.NewRecorder()

	handler.DeleteUser(rr, req)

	res := rr.Result()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}
