package main

import (
	"fmt"
	"github.com/Heavyymir/CharData_Aggregator/commands"
	"github.com/Heavyymir/CharData_Aggregator/config"
	"github.com/Heavyymir/CharData_Aggregator/internal/api"
	"strings"
	"github.com/chzyer/readline"
)

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	fields := strings.Fields(lowerCase)
	return fields
}


// REPL for command line interaction
func startRepl() {
	cfg := config.Config{
		CharDataClient: api.NewClient(),
	}

	// Initialise the completer to handle tab completion of internal commands
	completer := readline.NewPrefixCompleter(
		readline.PcItem("help"),
		readline.PcItem("select"),
		readline.PcItem("fetch"),
		readline.PcItem("exit"),
	)

	// Start the REPL
	rl, err := readline.NewEx(&readline.Config{
		Prompt:			"CharData > ",
		AutoComplete:	completer, 
	})
	
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// Defer close of the initialised REPL loop
	defer rl.Close()

	// Initialise an infinate loop to continue the REPL until the Exit command is called
	for{
		text, err := rl.Readline()
		if err != nil {
			return
		}

		words := cleanInput(text)
		if len(words) == 0 {
			fmt.Println("Empty Command. Please enter a valid command. Use command `help` for a list of valid commands")
			continue
		}
	

		if words[0] == "exit" {
			return
		}

		command, exists := commands.GetCommands()[words[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.Callback(&cfg, words[1:]...); err != nil {
			fmt.Println(err)
		}
	}	
}
