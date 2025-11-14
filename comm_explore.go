package main

import (
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a location name")
	} else if len(args) > 1 {
		return fmt.Errorf("location names with spaces are not supported yet")
	}

	PokeResp, err := cfg.pokeClient.ListPokemons(args[0])
	if err != nil {
		return err
	}

	for _, encounter := range PokeResp.PokemonEncounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}

	return nil
}
