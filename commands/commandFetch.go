package commands

import (
    "fmt"

    "github.com/Heavyymir/CharData_Aggregator/config"
    "github.com/Heavyymir/CharData_Aggregator/internal/api"
    "github.com/Heavyymir/CharData_Aggregator/internal/parsers/bbcf"
    //"github.com/Heavyymir/CharData_Aggregator/internal/parsers/ggst"
    //"github.com/Heavyymir/CharData_Aggregator/internal/parsers/sf6"
    //"github.com/Heavyymir/CharData_Aggregator/internal/parsers/sf3s"
)


// Command to Fetch character data once accessing a wiki
func commandFetch(cfg *config.Config, args ...string) error {

	// Ensure the entered strings aren't empty
	if cfg.Wiki.Name == "" || cfg.Game.Name == "" {
		return fmt.Errorf("select a wiki and game first")
	}
	
	if len(args) != 1 {
		return fmt.Errorf("usage: fetch <character>")
	}

	characterName := args[0]

	// Build the URL for the request
	requestURL, err := api.CharacterURL(
		cfg.Wiki,
		cfg.Game,
		characterName,
	)
	if err != nil {
		return err
	}

	fmt.Printf("Requesting: %s\n", requestURL)

	// Get the URL for Parsing
	data, err :=  cfg.CharDataClient.Fetch(requestURL)
	if err != nil {
		return err
	}

	fmt.Printf("Fetched %d bytes\n", len(data))
	
	// Switch case to handle parsers based on user inputed game
	switch cfg.Game.Slug {
		case "bbcf":
			fmt.Println("using bbcf parser")
			moves, err := bbcf.BBCFCharPageParser(data)
			if err != nil {
				return err
			}
			
			fmt.Printf("Character: %s\n", characterName)

			printMoveTable(moves)
				
			if err := writeMovesJSON(characterName, moves); err != nil {
				return err
			}
			
		default:
			return fmt.Errorf("no parser available for game: %s", cfg.Game.Name)
	}
	
	// Return request data to the user so they can see it has been successful 
	fmt.Printf("fetched %d bytes from %s\n", len(data), requestURL)

	return nil
}

