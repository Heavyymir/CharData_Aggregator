package sqlite

import(
        "database/sql"
        "fmt"

        "github.com/Heavyymir/CharData_Aggregator/internal/models"
)

// Function to retrieve all of a Characters Moves from internal SQL DB
func GetMoves(db *sql.DB, characterID int64) ([]models.Move, error) {
	// Query Data base, save to rows
	rows, err := db.Query(`
		SELECT id, input, name
		FROM moves
		WHERE character_id = ?
		ORDER BY id
		`, characterID)

	if err != nil {
		return nil, fmt.Errorf("get moves: %w", err)
	}

	defer rows.Close()

	// Initialise Slice to hold moves collected from SQL table
	var moves []models.Move

	// Iterate through rows
	for rows.Next() {
		var move models.Move
		var	moveID int64

		// Assign values to internal move struct
		if err := rows.Scan(
			&moveID,
			&move.Input,
			&move.Name,
			); err != nil {
				return nil, fmt.Errorf("scan move: %w", err)
			}

			// Use Framedata helper to Query and collect framedata information for Moves 
			move.FrameData, err = GetMoveFrameData(db, moveID)
			if err != nil {
				return nil, err
			}

			// Append move struct to moves
			moves = append(moves, move)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate moves: %w", err)
	}
	
	return moves, nil
}

// Function to collect data for an individual move
func GetMove(db *sql.DB, moveID int64) (models.Move, error) {
	// Intialise Move struct
	var move models.Move

	// Query Moves table, assign element values with Scan()
	err := db.QueryRow(`
		SELECT input, name
		FROM moves
		WHERE id = ?
	`, moveID).Scan(
		&move.Input, 
		&move.Name,
	) 

	if err != nil {
		return models.Move{}, fmt.Errorf("get move: %w", err)
	}

	// call framedata helper to query tables and assign framedata to move struct
	move.FrameData, err = GetMoveFrameData(db, moveID)
	if err != nil {
		return models.Move{}, fmt.Errorf("get move frame data: %w", err)
	}

	return move, nil
}

// Function to query and collect framedata for sql tables
func GetMoveFrameData(db *sql.DB, moveID int64) (map[string]models.Cell, error) {
	// Query framedata table, assign to rows
	rows, err := db.Query(`
		SELECT property, value
		FROM frame_data
		WHERE move_id = ?
	`, moveID)

	if err != nil {
		return nil, fmt.Errorf("get frame data: %w", err) 
	}
	
	defer rows.Close()

	// Intialise framedata map to match Move struct elements
	frameData := make(map[string]models.Cell)

	// Iterate over queried data in rows
	for rows.Next() {
		var property, value string

		// Use Scan to assign values to the framedata map on each pass
		if err := rows.Scan(
			&property,
			&value,
		); err != nil {
			return nil, fmt.Errorf("scan frame data: %w", err)
		}

		// Assign cell value to the framedata property key
		frameData[property] = models.Cell{Value: value}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate frame data: %w", err)
	}

	return frameData, nil
}
