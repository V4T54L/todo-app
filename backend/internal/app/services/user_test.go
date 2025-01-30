package services

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"todo_app_backend/internal/app/models"
)

// mockUserRepo is a mock implementation of repositories.UserRepoInterface.
type mockUserRepo struct {
	mock.Mock
}

func (m *mockUserRepo) CreateUser(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *mockUserRepo) DeleteUser(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockUserRepo) GetAllUsers(ctx context.Context) ([]models.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *mockUserRepo) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.User), args.Error(1)
}

// TestCreateUser tests the CreateUser method of UserService.
func TestCreateUser(t *testing.T) {
	mockRepo := new(mockUserRepo)
	userService := NewUserService(mockRepo)

	tests := []struct {
		name          string
		inputName     string
		inputEmail    string
		inputPassword string
		expectedError bool // New field for expected errors
		expectedUser  *models.User
	}{
		{"Valid Input", "John Doe", "john@example.com", "password123", false, &models.User{ID: 1, Name: "John Doe", Email: "john@example.com"}},
		{"Missing Name", "", "john@example.com", "password123", true, nil},
		{"Missing Email", "John Doe", "", "password123", true, nil},
		{"Missing Password", "John Doe", "john@example.com", "", true, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.expectedError {
				mockRepo.On("CreateUser", mock.Anything, mock.Anything).Return(nil).Once()
			}

			user, err := userService.CreateUser(context.Background(), tt.inputName, tt.inputEmail, tt.inputPassword)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, user) // user should be nil on error
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.expectedUser.Name, user.Name)
				assert.Equal(t, tt.expectedUser.Email, user.Email)
				assert.Empty(t, user.Password) // Password should be empty
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

// TestGetUserByID tests the GetUserByID method of UserService.
func TestGetUserByID(t *testing.T) {
	mockRepo := new(mockUserRepo)
	userService := NewUserService(mockRepo)

	tests := []struct {
		name         string
		userID       int
		mockUser     *models.User
		mockError    error
		expectedUser *models.User
	}{
		{"User Found", 1, &models.User{ID: 1, Name: "John Doe"}, nil, &models.User{ID: 1, Name: "John Doe"}},
		{"User Not Found", 2, nil, errors.New("user not found"), nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("GetUserByID", mock.Anything, tt.userID).Return(tt.mockUser, tt.mockError).Once()

			user, err := userService.GetUserByID(context.Background(), tt.userID)

			if tt.expectedUser == nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedUser.ID, user.ID)
				assert.Equal(t, tt.expectedUser.Name, user.Name)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

// TestGetAllUsers tests the GetAllUsers method of UserService.
func TestGetAllUsers(t *testing.T) {
	mockRepo := new(mockUserRepo)
	userService := NewUserService(mockRepo)

	tests := []struct {
		name          string
		mockUsers     []models.User
		mockError     error
		expectedUsers []models.User
	}{
		{"Users Found", []models.User{{ID: 1, Name: "John Doe"}}, nil, []models.User{{ID: 1, Name: "John Doe"}}},
		{"No Users", []models.User{}, nil, []models.User{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("GetAllUsers", mock.Anything).Return(tt.mockUsers, tt.mockError).Once()

			users, err := userService.GetAllUsers(context.Background())

			if tt.mockError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(tt.expectedUsers), len(users))
				assert.ElementsMatch(t, tt.expectedUsers, users)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

// TestDeleteUser tests the DeleteUser method of UserService.
func TestDeleteUser(t *testing.T) {
	mockRepo := new(mockUserRepo)
	userService := NewUserService(mockRepo)

	tests := []struct {
		name      string
		userID    int
		mockError error
	}{
		{"User Deleted", 1, nil},
		{"User Not Found", 2, errors.New("could not delete user")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("DeleteUser", mock.Anything, tt.userID).Return(tt.mockError).Once()

			err := userService.DeleteUser(context.Background(), tt.userID)

			if tt.mockError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
