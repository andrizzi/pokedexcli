package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/andrizzi/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeClient       pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]

		cmd, exists := getCommands()[cmdName]
		if !exists {
			fmt.Printf("Unknown command: %s\n", cmdName)
			continue
		}
		err := cmd.callback(cfg)
		if err != nil {
			fmt.Printf("Error executing command %s: %v\n", cmdName, err)
		}

	}
}

func cleanInput(input string) []string {
	output := strings.ToLower(input)
	return strings.Fields(output)
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the map of the Pokemon world",
			callback:    commandMap,
		},
	}
}
