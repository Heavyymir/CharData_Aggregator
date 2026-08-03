package commands

import(
	"errors"
	"fmt"

	"github.com/Heavyymir/CharData_Aggregator/config"
	"github.com/Heavyymir/CharData_Aggregator/catalog"
)


// Command to allow users to walk through selections of Wikis, Games and Characters
func commandSelect(cfg *config.Config, args ...string) error {
	switch {
		// Base case if a user attempts a wildcard search to see which wikis are available
		case len(args) == 1 && args[0] == "*":
			printWikis()
			return nil

		// Case to check for the games under a wiki that the program can handle.
		case len(args) == 2 && args[1] == "*":
			wiki, exists := catalog.Wikis[args[0]]
			if !exists {
				return errors.New("unknown wiki: %s", args[0])
			}

			game, exists := wiki.Games[args[1]]
			if !exists{
				return fmt.Errorf("unknown game: %s", args[1])
			}		

			// Store the selected wiki and game in config
			cfg.Wiki = wiki
			cfg.Game = game

			// Show the User their selected Wiki and Game.
			fmt.Printf("selected %s: %s\n", wiki.Name, game.Name)

			return nil

		default:
			return errors.New("usage: select * | select <wikiname> *")
	}
	return nil
}

// Wrapper to print wikis to console
func printWikis() {
	for key, wiki := range catalog.Wikis {
		fmt.Printf("%s: %s\n", key, wiki.Name)
	}
}

// Wrapper to print available games to console
func printGames(wiki catalog.Wiki) {
	for key, game := range wiki.Games {
		fmt.Printf("%s: %s\n", key, game.Name)
	}
}
