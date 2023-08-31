package main

import (
	"errors"
	"fmt"
)

func mapN(cfg *config) error {
	locResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocURL)
	if err != nil {
		return err
	}
	cfg.nextLocURL = locResp.Next
	cfg.prevLocURL = locResp.Previous
	for _, loc := range locResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func mapB(cfg *config) error {
	if cfg.prevLocURL == nil {
		return errors.New("You are on the First page (use the 'map' command)")
	}

	locResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocURL)
	if err != nil {
		return nil
	}

	cfg.nextLocURL = locResp.Next
	cfg.prevLocURL = locResp.Previous

	for _, loc := range locResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
