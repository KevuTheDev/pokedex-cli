package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return fmt.Errorf("pokemon is yet to be caught")
	}

	fmt.Println("Pokemon in Pokedex:")

	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
