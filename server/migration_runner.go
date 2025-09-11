package server

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"
)

func RunMigrations(db *sql.DB, migrationDir string) error {

	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS migration_history (
            id SERIAL PRIMARY KEY,
            filename VARCHAR(255) UNIQUE NOT NULL,
            executed_at TIMESTAMP DEFAULT NOW()
        )
    `)
	if err != nil {
		return fmt.Errorf("failed to create migration_history table: %w", err)
	}

	files, err := filepath.Glob(filepath.Join(migrationDir, "*.sql"))
	if err != nil {
		return fmt.Errorf("failed to read migration files: %w", err)
	}

	sort.Strings(files)

	for _, file := range files {
		filename := filepath.Base(file)

		var exists bool
		err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM migration_history WHERE filename=$1)", filename).Scan(&exists)
		if err != nil {
			return fmt.Errorf("failed to check migration history: %w", err)
		}

		if exists {
			log.Printf("Skipping migration (already executed): %s", filename)
			continue
		}

		content, err := ioutil.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", file, err)
		}

		// Split query dengan delimiter ;
		queries := strings.Split(string(content), ";")
		for _, query := range queries {
			query = strings.TrimSpace(query)
			if query == "" {
				continue
			}
			_, err := db.Exec(query)
			if err != nil {
				return fmt.Errorf("failed to execute query from %s: %w", file, err)
			}
		}

		_, err = db.Exec("INSERT INTO migration_history (filename) VALUES ($1)", filename)
		if err != nil {
			return fmt.Errorf("failed to insert migration history for %s: %w", filename, err)
		}

		log.Printf("Migration executed successfully: %s", filename)
	}

	log.Println("All migrations executed successfully")
	return nil
}
