package main

import (
	"fmt"
)

func commandInspect(c *config, pokemonNames []string) error {
	name := pokemonNames[0]

	if pokemon, ok := c.pokedex[name]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Print("Stats:\n")
		for _, s := range pokemon.Stats {
			fmt.Printf(" -%s: %d\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Print("Types:\n")
		for _, t := range pokemon.Types {
			fmt.Printf(" -%s: %d\n", t.Type.Name, t.Slot)
		}
		fmt.Print("\n")
	} else {
		fmt.Print("You haven't caught this Pokemon, yet\n\n")
	}

	return nil
}

