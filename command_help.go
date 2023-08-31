package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex help")
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println()
	for _, cmd := range Commands() {
		fmt.Printf("%s - %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
