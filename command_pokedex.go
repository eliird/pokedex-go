package main

import (
	"errors"
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("no pokemon have been captured yet")
	}
	fmt.Println("Pokemon available in pokedex: ")
	count := 0
	for k, _ := range cfg.caughtPokemon {
		fmt.Printf("%d : %s\n", count, k)
		count++
	}
	return nil
}
