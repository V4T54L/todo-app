package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo_app_backend/internal/app/models"
	"todo_app_backend/internal/app/services"

	"github.com/go-chi/chi/v5"
)

type TodoHandler struct {
	Service services.TodoServiceInterface
}

// TODO : Fetch userId from request
const (
	userId = 1
)

// NewTodoHandler initializes a new TodoHandler.
func NewTodoHandler(service services.TodoServiceInterface) *TodoHandler {
	return &TodoHandler{Service: service}
}

// CreateTodo handles the creation of a new todo.
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdTodo, err := h.Service.CreateTodo(r.Context(), userId, todo.Title, todo.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTodo)
}

// GetAllTodos retrieves all todos.
func (h *TodoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.Service.GetAllTodos(r.Context(), userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

// GetTodoByID retrieves a todo by ID.
func (h *TodoHandler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	fmt.Println("id : ", id)

	todo, err := h.Service.GetTodoByID(r.Context(), userId, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

// UpdateTodo updates a todo by ID.
func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateTodo(r.Context(), userId, id, todo.Title, todo.Content, todo.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteTodo deletes a todo by ID.
func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteTodo(r.Context(), userId, id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

var _ TodoHandlerInterface = (*TodoHandler)(nil)
