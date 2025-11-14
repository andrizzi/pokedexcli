package main

import "fmt"

func commandInspect(cfg *config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide exactly one Pokemon name to inspect")
	}
	pokemonName := args[0]

	pokemon, caught := cfg.caughtPokemon[pokemonName]
	if !caught {
		return fmt.Errorf("you have not caught a Pokemon named %s", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
