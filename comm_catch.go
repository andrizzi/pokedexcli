package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a Pokemon name")
	} else if len(args) > 1 {
		return fmt.Errorf("pokemon names are just one word")
	}

	PokeResp, err := cfg.pokeClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s... ", PokeResp.Name)
	// wait a bit to simulate the throw
	time.Sleep(2 * time.Second)
	// use the pokemon's "base experience" to determine the chance of catching it
	captureChance := PokeResp.BaseExperience
	if captureChance > 255 {
		captureChance = 255
	}
	roll := rand.Intn(256) // 0-255
	if roll < captureChance {
		fmt.Printf("Gotcha! %s was caught!\n", PokeResp.Name)
		cfg.caughtPokemon[PokeResp.Name] = PokeResp
	} else {
		fmt.Printf("%s escaped!\n", PokeResp.Name)
	}

	return nil
}
