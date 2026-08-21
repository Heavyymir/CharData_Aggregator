package commands

import(
	"fmt"
	"strings"

	"github.com/Heavyymir/CharData_Aggregator/config"
	"github.com/Heavyymir/CharData_Aggregator/internal/storage/sqlite"
)

// Cli Command to pull framedata for a character 
func commandFrames(cfg *config.Config, args ...string) error {
	// Verify arguments
	if len(args) != 1 {
		return fmt.Errorf("usage frames <character_name>")
	}

	if cfg.Game.Name == "" {
			return fmt.Errorf("select a game first")
		}

	// Assign Character slug
	characterSlug := args[0]

	// Get character ID from SQL table
	characterID, err := sqlite.GetCharacterID(
		cfg.DB,
		cfg.Game.Slug,
		characterSlug,
	)
	if err != nil {
		return err
	}

	// Get moves from SQL table, print framedata to console
	moves, err := sqlite.GetMoves(cfg.DB, characterID)
	if err != nil {
		return err
	}

	// Debug/Output
	fmt.Printf("moves loaded: %d\n", len(moves))


	// Print logic for SQL framedata grids
	for _, move := range moves {
		fmt.Printf("=== %s (%s) ===\n", move.Name, move.Input)
		
		for i, grid := range move.FrameDataGrids {
			if i == 0 {
				fmt.Println("[Base Frame Data]")
			} else {
				fmt.Printf("[Additional Data - Grid %d]\n", i)
			}
			printGrid(grid)
			fmt.Println()
		}
			
		if len(move.Notes) > 0 {
			fmt.Println("Notes:")
			for _, note := range move.Notes {
				fmt.Printf("  • %s\n", note)
			}
			fmt.Println()
		}

		fmt.Println(strings.Repeat("=", 60))
	}

	return nil
	
}
