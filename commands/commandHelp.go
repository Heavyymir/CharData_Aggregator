package commands

import (
	"fmt"

	"github.com/Heavyymir/CharData_Aggregator/config"
)

func commandHelp(cfg *config.Config, args ...string) error {
	fmt.Println(`Welcome to the Character Data Aggregator

Usage:
help: Displays a list of commands usable in the Character Data Aggregator
exit: Exits the tool
select: Allows user to select a wiki and game. Usage: select * | select <wiki> * | select <wiki> <game>
fetch: Fetches a character page for a game. Usage: fetch <character>
discover: Discovers character names for a game, and saves to a local json.`)

	return nil
}
