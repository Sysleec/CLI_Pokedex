package main

import (
	"time"

	"github.com/Sysleec/CLI_Pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	Repl(cfg)
}
