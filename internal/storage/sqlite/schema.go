package sqlite


import (
	"database/sql"
	"fmt"
)

func InitialiseSchema(db *sql.DB) error {
// Define an imutatble internal SQL schema to create required tables if they do not already exist.
	const schema = `
CREATE TABLE IF NOT EXISTS characters (
    id INTEGER PRIMARY KEY,
    game TEXT NOT NULL,
    name TEXT NOT NULL,
    slug TEXT NOT NULL,
    source_url TEXT,
    UNIQUE(game, slug)
);

CREATE TABLE IF NOT EXISTS moves (
    id INTEGER PRIMARY KEY,
    character_id INTEGER NOT NULL,
    input TEXT,
    name TEXT NOT NULL,
    FOREIGN KEY (character_id) REFERENCES characters(id)
);

CREATE TABLE IF NOT EXISTS frame_data (
    id INTEGER PRIMARY KEY,
    move_id INTEGER NOT NULL,
    property TEXT NOT NULL,
    value TEXT,
    tooltip TEXT,
    position INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY (move_id) REFERENCES moves(id),
    UNIQUE(move_id, property)
);
`
	if _, err := db.Exec(schema); err != nil {
		return fmt.Errorf("initialised schema: %w", err)
	}
	return nil
}
