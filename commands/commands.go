package commands

type cliCommand struct (
	name			string
	description		string
	callback		func(config, ...string) error	
)

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name:			"exit",
			description:	"command to exit CharData_Aggregator",
			callback:		commandExit,
		},
		"help": {
			name:			"help",
			description:	"displays information on commands usable in CharData_Aggregator",
			callback:		commandHelp,
		},
	}
}
