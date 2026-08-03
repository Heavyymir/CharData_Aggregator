package commands

import "fmt"

func commandHelp(config, args ...string) error {
	fmt.Println(`Welcome to the Character Data Aggregator

Usage:
help: Displays a list of commands usable in the Character Data Aggregator
exit: Exits the tool`)

	return nil
}
