package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo_app_backend/api"
	v1 "todo_app_backend/api/v1"
	"todo_app_backend/config"
	"todo_app_backend/internal/app/repositories"
	"todo_app_backend/internal/app/services"
	"todo_app_backend/internal/database"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	done <- true
}

func main() {
	if os.Getenv("ENVIRONMENT") != "PRODUCTION" {
		err := config.LoadConfigurationFile(".env")
		if err != nil {
			log.Fatal("Error loading .env : ", err)
		}
	}
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal("Error loading config : ", err)
	}

	db, err := database.NewPostgreSQLDB(cfg.DatabaseURI, cfg.MaxIdleConns, cfg.MaxOpenConns)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	userHandler := v1.NewUserHandler(services.NewUserService(repositories.NewUserRepository(db.GetConn())))
	todoHandler := v1.NewTodoHandler(services.NewTodoService(repositories.NewTodoRepository(db.GetConn())))

	server := http.Server{
		Addr:         ":8080",
		Handler:      api.SetupRouter(userHandler, todoHandler),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  time.Minute,
	}

	done := make(chan bool, 1)

	go gracefulShutdown(&server, done)

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("Graceful shutdown complete.")
}
