package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg config) {
	for {
		fmt.Print("Pokedex > ")
		newScanner := bufio.NewScanner(os.Stdin)
		scanResult := newScanner.Scan()

		if newScanner.Text() == "" {
			fmt.Print("Please, provide a valid command\n\n")
			continue
		}

		if scanResult {
			command, predicates := cleanInput(newScanner.Text())
			if c, ok := getCommands()[command]; !ok {
				fmt.Print("Command was not found\n\n")
				continue
			} else {
				c.callback(&cfg, predicates)
			}
		} 
	}
}

func cleanInput(t string) (string, []string) {
	output := strings.ToLower(t) 
	words := strings.Fields(output)
	return words[0], words[1:]
}

func getCommands() commands {
	return commands {
		"help": {
			name:		"help",	
			description: 	"Displays a help message",	
			callback:	commandHelp,	
		},
		"exit": {
			name:		"exit",	
			description: 	"Exit the Pokedex",	
			callback: 	commandExit,	
		},
		"map": {
			name:		"map",	
			description: 	"Show next 20 Pokemon world locaion names",	
			callback: 	commandMap,	
		},
		"mapb": {
			name:		"mapb",	
			description: 	"Show previous 20 Pokemon world locaion names",	
			callback: 	commandMapb,	
		},
		"explore": {
			name:		"explore",
			description:	"See which Pokemons live in the provided area",
			callback:	commandExplore,
		},
		"catch": {
			name:		"catch",
			description:	"Try to catch the chosen Pokemon into your Pokedex",
			callback:	commandCatch,
		},
		"inspect": {
			name:		"inspect",
			description:	"Inspect an already caught Pokemon",
			callback:	commandInspect,
		},
		"pokedex": {
			name:		"pokedex",
			description:	"List all the Pokemons in the Pokedex",
			callback:	commandPokedex,
		},
	}
}

