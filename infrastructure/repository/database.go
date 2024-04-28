package repository

import (
	"database/sql"
	"log"
)

func InitializeSqlLiteDB(db *sql.DB) {
	createMigrationsTable(db)

	migrations := map[string]func(*sql.DB) error{
		"M001CreatePomodoroTable": func(db *sql.DB) error {
			_, err := db.Exec("CREATE TABLE pomodoro (id TEXT PRIMARY KEY, duration INTEGER, status TEXT)")
			return err
		},
	}

	for name, migration := range migrations {
		if !isMigrationApplied(db, name) {
			err := migration(db)
			if err != nil {
				log.Fatalf("Error applying migration %v: %v", name, err)
			}

			_, err = db.Exec("INSERT INTO migrations (name) VALUES (?)", name)
			if err != nil {
				log.Fatalf("Error saving migration %v: %v", name, err)
			}
		}
	}
}

func isMigrationApplied(db *sql.DB, s string) bool {
	var name string
	err := db.QueryRow("SELECT name FROM migrations WHERE name = ?", s).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Fatalf("Error checking if migration %v is applied: %v", s, err)
	}

	return true
}

func createMigrationsTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS migrations (name TEXT PRIMARY KEY)")
	if err != nil {
		log.Fatalf("Error creating migrations table: %v", err)
	}
}
