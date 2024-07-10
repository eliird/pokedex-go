package main

import (
	"time"

	"github.com/eliird/pokedex-go/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	currentPageURL          *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Minute * 5),
	}
	startRepl(&cfg)

}
