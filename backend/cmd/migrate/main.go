package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"todo_app_backend/config"
	"todo_app_backend/internal/database"
)

// executeMigration runs the SQL commands from the specified file
func executeMigration(db *sql.DB, filePath string) error {
	// Read the migration file
	sqlFile, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	// Execute the migration
	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	return nil
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

	conn := db.GetConn()

	// Usage: Migrate up or down based on command line argument
	if len(os.Args) < 2 {
		log.Fatalf("No argument provided, expected [up|down]")
	}

	action := os.Args[1]
	switch action {
	case "up":
		if err := executeMigration(conn, "db/migrations/001_init.up.sql"); err != nil {
			log.Fatalf("Migration up failed: %v", err)
		}
		log.Println("Migration up applied successfully")
	case "down":
		if err := executeMigration(conn, "db/migrations/001_init.down.sql"); err != nil {
			log.Fatalf("Migration down failed: %v", err)
		}
		log.Println("Migration down applied successfully")
	case "seed":
		if err := executeMigration(conn, "db/seeders/001_init.sql"); err != nil {
			log.Fatalf("Migration seed failed: %v", err)
		}
		log.Println("Migration down applied successfully")
	default:
		log.Fatalf("Unknown command: %s. Use 'up' or 'down'.", action)
	}
}
