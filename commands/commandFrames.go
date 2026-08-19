package commands

import(
	"fmt"

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

	for _, move := range moves {
		fmt.Printf("%s %s\n", move.Name, move.Input)
		fmt.Println("--------------------------------------------------")
		for _, header := range move.Headers {
			fmt.Printf("%s: %s\n", header, move.FrameData[header].Value)
		}
	}
	return nil
}
