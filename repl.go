package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Turns off the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Search",
			callback:    callbackMap,
		},
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex-cli > ")
		scanner.Scan()

		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("invalid command")
			fmt.Println()
			continue
		}

		command.callback()
		fmt.Println()

	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
