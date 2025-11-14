package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("No caught Pokemon yet.")
		return nil
	}

	fmt.Println("Caught Pokemon:")
	for name := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", name)
	}
	return nil
}
