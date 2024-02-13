package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(c *config, pokemonNames []string) error {
	name := pokemonNames[0]
	pokemon, err := c.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	successfullyCaught := pokemon.BaseExperience / (r.Intn(pokemon.BaseExperience) + 1) > 3 

	if successfullyCaught {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Print("You may now inspect it with the inspect command.\n\n")
		c.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n\n", pokemon.Name)
	}

	return nil
}
