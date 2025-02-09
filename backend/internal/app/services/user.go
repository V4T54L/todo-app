package services

import (
	"context"
	"errors"
	"todo_app_backend/internal/app/models"
	"todo_app_backend/internal/app/repositories"
	"todo_app_backend/internal/app/utils"
)

// UserService defines methods related to user operations.
type UserService struct {
	UserRepo repositories.UserRepoInterface
	// hashFunc func(string) string
}

// NewUserService initializes a new UserService.
func NewUserService(userRepo repositories.UserRepoInterface) *UserService {
	return &UserService{UserRepo: userRepo}
}

// CreateUser creates a new user with hashed password.
func (s *UserService) CreateUser(ctx context.Context, name, email, password string) (*models.User, error) {
	// Validate input
	if name == "" || email == "" || password == "" {
		return nil, errors.New("name, email, and password are required")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create user model
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	// Store user in DB
	if err := s.UserRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	user.Password = ""

	return user, nil
}

// GetUserByID retrieves a user by ID.
func (s *UserService) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	user, err := s.UserRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByCreds retrieves a user by Email and Password.
func (s *UserService) GetUserByCreds(ctx context.Context, email, password string) (*models.User, error) {
	user, err := s.UserRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	utils.CheckPasswordHash(password, user.Password)
	user.Password = ""
	return user, nil
}

// GetAllUsers retrieves all users.
func (s *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.UserRepo.GetAllUsers(ctx)
}

// DeleteUser deletes a user by ID.
func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.UserRepo.DeleteUser(ctx, id)
}

var _ UserServiceInterface = (*UserService)(nil)
