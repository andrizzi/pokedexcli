package main

import (
	"time"

	"github.com/andrizzi/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeClient: pokeClient,
	}

	startRepl(cfg)
}
