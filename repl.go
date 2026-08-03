package main

import (
	"bufio"
	"fmt"
	"github.com/Heavyymir/CharData_Aggregator/commands"
	"github.com/Heavyymir/CharData_Aggregator/config"
	"github.com/Heavyymir/CharData_Aggregator/internal/api"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	fields := strings.Fields(lowerCase)
	return fields
}

func startRepl() {
	cfg := config.Config{
		CharDataClient: api.NewClient(),
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("CharData > ")
		if !scanner.Scan() {
			return
		}
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			fmt.Println("Empty Command. Please use a valid command. Type 'help' for more information.")
			continue
		}
		command, exists := commands.GetCommands()[words[0]]
		if exists {
			args := words[1:]
			err := command.Callback(&cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command, please refer to 'help' for more information")
		}
	}
	return
}
