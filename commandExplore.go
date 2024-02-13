package main

import "fmt"

func commandExplore(c *config, areas []string) error {
	area := areas[0]

	fmt.Printf("Exploring %s...\n", area)
	ap, err := c.pokeapiClient.GetPokemons(&area) 
	if err != nil {
		return err
	}

	fmt.Print("Found Pokemon:\n")
	for _, pokemon := range ap.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
		
	fmt.Print("\n")

	return nil
}
