package main

import "fmt"

func commandPokedex(c *config, _ []string) error {
	fmt.Print("Your Pokedex:\n")
	for k := range c.pokedex {
		fmt.Printf(" - %s\n", k)	
	} 
	fmt.Print("\n")

	return nil
}

