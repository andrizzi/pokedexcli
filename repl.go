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
	caughtPokemon    map[string]pokeapi.RespDetailedPokemons
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]
		cmdArg := words[1:]

		cmd, exists := getCommands()[cmdName]
		if !exists {
			fmt.Printf("Unknown command: %s\n", cmdName)
			continue
		}
		err := cmd.callback(cfg, cmdArg)
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
	callback    func(*config, []string) error
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
		"mapb": {
			name:        "mapb",
			description: "Go back to the previous map page",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location by its name",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Catch a Pokemon by its name or ID",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Inspect a caught Pokemon by its name",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught Pokemon",
			callback:    commandPokedex,
		},
	}
}
