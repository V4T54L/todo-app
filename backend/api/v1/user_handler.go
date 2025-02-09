package v1

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"todo_app_backend/internal/app/models"
	"todo_app_backend/internal/app/services"
	"todo_app_backend/internal/app/utils"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	UserService services.UserServiceInterface
}

// NewUserHandler initializes a new UserHandler.
func NewUserHandler(userService services.UserServiceInterface) *UserHandler {
	return &UserHandler{UserService: userService}
}

// CreateUser handles the user registration request.
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.User

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.UserService.CreateUser(r.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUserByID handles retrieving a user by ID.
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.UserService.GetUserByID(r.Context(), id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// LoginHandler handles Login request.
func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.User

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.UserService.GetUserByCreds(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	tokenDetails := models.Token{UserID: user.ID, Exp: time.Now().Add(time.Minute * 15)}
	token, err := utils.GenerateToken(tokenDetails)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
		"user":  user,
	})
}

// GetAllUsers handles retrieving all users.
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserService.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// DeleteUser handles deleting a user by ID.
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.UserService.DeleteUser(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}

var _ UserHandlerInterface = (*UserHandler)(nil)
