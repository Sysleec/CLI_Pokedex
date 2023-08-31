package main

import (
	"errors"
	"fmt"
)

func Explore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide the location name (Example 'explore lake-valor-area')")
	}

	name := args[0]
	loc, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	PokCount := 0
	fmt.Printf("Exploring location '%s'...\n", loc.Name)
	fmt.Println("Found Pokemon: ")
	for _, pok := range loc.PokemonEncounters {
		fmt.Printf(" - %s\n", pok.Pokemon.Name)
		PokCount++
	}
	fmt.Printf("Total - %v\n", PokCount)
	return nil
}
