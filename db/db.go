package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	// Default to local, override with env var for production
	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		dataDir = "./data" // Local development default
	}

	if err := os.MkdirAll(dataDir, 0755); err != nil {
		log.Printf("Warning: Could not create data directory: %v", err)
	}

	dbPath := filepath.Join(dataDir, "api.db")

	DB, err = sql.Open("sqlite3", dbPath)

	if err != nil {
		panic("could not connect to the database")
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		log.Fatal("Could not ping database:", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	log.Printf("Database connected successfully at %s", dbPath)

	createTables()
}

func createTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("could not create users table")
	}

	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT NOT NULL,
      description TEXT NOT NULL,
      location TEXT NOT NULL,
      date_time DATETIME NOT NULL,
      user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id)
    )
  `

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("could not create events table")
	}

	createRegistrationsTable := `
		CREATE TABLE IF NOT EXISTS registrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY(event_id) REFERENCES events(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("could not create registrations table")
	}

	log.Println("All tables created successfully")
}
