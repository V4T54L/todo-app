package api

import (
	"net/http"
	"time"
	v1 "todo_app_backend/api/v1"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// SetupRouter initializes the API routes.
func SetupRouter(userHandler v1.UserHandlerInterface, todoHandler v1.TodoHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(RespondJsonMiddleware)
	// r.Use()

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/v1", func(r chi.Router) {

		// user routes
		r.Route("/users", func(r chi.Router) {
			r.Post("/", userHandler.CreateUser)
			r.Get("/", userHandler.GetAllUsers)
			r.Get("/{id}", userHandler.GetUserByID)
			r.Delete("/{id}", userHandler.DeleteUser)
		})

		// todo routes
		r.Route("/todos", func(r chi.Router) {
			r.Post("/", todoHandler.CreateTodo)
			r.Get("/", todoHandler.GetAllTodos)
			r.Get("/{id}", todoHandler.GetTodoByID)
			r.Delete("/{id}", todoHandler.DeleteTodo)
			r.Put("/{id}", todoHandler.UpdateTodo)
		})

	})

	return r
}
