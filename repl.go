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
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch {pokemon-name}",
			description: "Capture a pokemon",
			callback:    callbackCatch,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the Pokedex",
			callback:    callbackExit,
		},
		"explore": {
			name:        "explore {location-area}",
			description: "Pokemon's in an area",
			callback:    callbackExplore,
		},
		"inspect": {
			name:        "inspect {pokemon-name}",
			description: "Inspect caught pokemon",
			callback:    callbackInspect,
		},
		"map": {
			name:        "map",
			description: "Search",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Search",
			callback:    callbackMapb,
		},
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
	}
}

func startRepl(cfg *config) {
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
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("invalid command")
			fmt.Println()
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println()
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
