package main

import "github.com/Sysleec/CLI_Pokedex/internal/pokeapi"

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	nextLocURL    *string
	prevLocURL    *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func Commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    mapN,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    mapB,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    Explore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    Catch,
		},
	}
}
