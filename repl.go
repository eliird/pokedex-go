package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	availableCommands := getCommnads()

	for {
		fmt.Print("Please enter the name of the pokemon > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		// handle enter without data
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		command.callback()
	}

}

type cliCommnad struct {
	name        string
	description string
	callback    func()
}

func getCommnads() map[string]cliCommnad {

	return map[string]cliCommnad{
		"help": {
			name:        "help",
			description: "prints the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "turns off pokedex",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	//extract the pompts from the user
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered) //splits the sentence into the words
	return words
}
