package commands

import (
    "fmt"

    "github.com/Heavyymir/CharData_Aggregator/config"
    "github.com/Heavyymir/CharData_Aggregator/internal/api"
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

	data, err :=  cfg.CharDataClient.Fetch(requestURL)
	if err != nil {
		return err
	}

	// Return request data to the user so they can see it has been successful 
	fmt.Printf("Fetched %d bytes from %s\n", len(data), requestURL)

	return nil
}
