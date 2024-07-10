package main

import (
	"fmt"
)

func callbackHelp() {
	fmt.Println("Welcome to the pokedex help menu!")
	fmt.Println("Here are your available commanads")
	availableCommands := getCommnads()

	for _, command := range availableCommands {
		fmt.Printf("- %s : %s\n", command.name, command.description)
	}
	fmt.Println("")
}