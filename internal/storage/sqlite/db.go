package sqlite

import(
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)


// Function to initialise a local SQL database for the storage of character data
func Open(path string) (*sql.DB, error) {
	// Open and prepare the SQL database
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	// Call Ping() to verify the connection to the opened database
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	// Call Exec() to run SQL without returning rows
	if _, err := db.Exec(`PRAGMA foreign_keys= ON`); err != nil {
		db.Close()
		return nil, fmt.Errorf("enable foreign keys: %w", err)
	}
	// Return the created database
	return db, nil	
}
