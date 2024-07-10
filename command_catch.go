package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	const threshhold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshhold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}
	fmt.Printf("%s was caught\n", pokemonName)
	cfg.caughtPokemon[pokemonName] = pokemon
	return nil
}
