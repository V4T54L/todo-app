package repositories

import (
	"database/sql"
	"errors"
	"todo_app_backend/internal/app/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
	return err
}

// GetUserByID retrieves a user by ID
func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, name, email, password, created_at FROM users WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return user, err
}

// GetAllUsers retrieves all users
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	query := `SELECT id, name, email, password, created_at FROM users`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// DeleteUser removes a user by ID
func (r *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
