package main

import "fmt"

func commandHelp(c *config, _ []string) error {
	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")
	for k, v := range getCommands() {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	fmt.Print("\n")
	return nil 
}
