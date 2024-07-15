# Pokedex CLI Game

A command-line interface (CLI) game that lets you explore, catch, and inspect Pokémon. The game features various commands to interact with a Pokémon-themed REPL (Read-Eval-Print Loop).

## Features

- **Help Command**: Lists all available commands within the program.
- **Map Commands**: Integrates with the Pokémon API to fetch and display a list of 20 locations to explore. Navigate using `mapn` (next) and `mapb` (back) commands.
- **Caching Mechanism**: Stores API call results for 5 minutes to enhance performance and reduce redundant requests. Automatically clears unused items after the specified duration.
- **Exploration**: Explore different areas using the `explore` command followed by the area's name.
- **Catching Pokémon**: Attempt to catch Pokémon with the `catch` command followed by the Pokémon's name. Success is determined by a random chance based on the Pokémon's base experience.
- **Pokedex Command**: View a list of all captured Pokémon.
- **Inspect Command**: Get detailed statistics of a specific Pokémon using the `inspect` command followed by the Pokémon's name.

## Installation

To install the Pokedex CLI Game, use the following command:

```sh
go install github.com/eliird/pokedex-go

$ pokedex-go
Welcome to the Pokedex CLI Game!

Type 'help' to see a list of commands.

> map
Fetching locations...
1. Pallet Town
2. Viridian City
3. Pewter City
...

> mapn
Moved to the next location.

> explore Pallet Town
Exploring Pallet Town...
You encountered a Pikachu!

> catch Pikachu
Attempting to catch Pikachu...
Success! You caught Pikachu.

> pokedex
Your captured Pokémon:
1. Pikachu

> inspect Pikachu
Name: Pikachu
Type: Electric
Base Experience: 112
...
