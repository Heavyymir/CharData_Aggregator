package commands

import "github.com/Heavyymir/CharData_Aggregator/config"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*config.Config, ...string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "command to exit CharData_Aggregator",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "displays information on commands usable in CharData_Aggregator",
			Callback:    commandHelp,
		},
		"select": {
			Name: 		 "select",
			Description: "selects a wiki and game",	
			Callback:	 commandSelect,
		},
		"fetch": {
			Name:		 "fetch",
			Description: "fetches character data from a wiki",
			Callback:	 commandFetch,
		},
	}
}
