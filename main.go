package main

import (
	"math/rand"
	"time"

	"github.com/andrizzi/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.RespDetailedPokemons{},
		pokeClient:    pokeClient,
	}

	rand.Seed(time.Now().UnixNano())

	startRepl(cfg)
}
