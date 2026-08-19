package sqlite


import(
        "database/sql"
        "fmt"

        "github.com/Heavyymir/CharData_Aggregator/internal/models"
)

// Function to save a characters moves to the DB
func SaveMoves(db *sql.DB, characterID int64, moves []models.Move) error {

	// Start the SQL transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin move transaction: %w", err)
	}
	
	defer tx.Rollback()

	// Loop over moves slice and call saveMove on each entry
	for _, move := range moves {
		if _, err := saveMove(tx, characterID, move); err != nil {
			return err
		}
	}

	// Once the loop completes, call .Commit() to finalise changes to the SQL table
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit moves: %w", err)
	}
	return nil
}
				
// Helper fo SaveMoves to handle the Insertion of individal moves into the DB 
func saveMove(tx *sql.Tx, characterID int64, move models.Move) (int64, error) {

	// Call .Exec() method on the input SQL transaction to insert data into DB
	result, err := tx.Exec(`
	INSERT INTO moves (character_id, input, name)
	VALUES (?, ?, ?)
	`, characterID, move.Input, move.Name)

	if err != nil {
		return 0, fmt.Errorf("save move: %w", err)
	}

	// Assign a Move ID to the insertion
	moveID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get move ID: %w", err)
	}

	seen := make(map[string]bool)

	// Insert headers to the framedata data using position to maintain ordering and match to data
	for position, property := range move.Headers {
		if seen[property] {
			return 0, fmt.Errorf(
				"duplicate header %q in move %q",
				property,
				move.Name,
			)
		}

		seen[property] = true

		fmt.Printf("inserting: move=%q id=%q property=%q",
			move.Name, moveID, property)

		cell, ok := move.FrameData[property]
		if !ok{
			continue
		}
					
		_, err := tx.Exec(`
		INSERT INTO frame_data
			(move_id, property, value, tooltip, position)
		VALUES (?, ?, ?, ?, ?)
		`,
		moveID,
		property,
		cell.Value,
		cell.Tooltip,
		position,
		)
		if err != nil {
			return 0, fmt.Errorf("save frame data: move=%q moveID=%q property=%q: %w", 
			move.Name, moveID, property, err,
			)
		}
	}

	return moveID, nil	
}
