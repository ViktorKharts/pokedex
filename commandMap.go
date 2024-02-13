package main

import (
	"errors"
	"fmt"
)

func commandMap(c *config, _ []string) error {
	locs, err := c.pokeapiClient.GetLocations(c.Next)
	if err != nil {
		return err
	}

	c.Next = locs.Next
	c.Previous = locs.Previous
	for _, l := range locs.Results {
		fmt.Println(l.Name)
	}
	fmt.Print("\n")

	return nil
}

func commandMapb(c *config, _ []string) error {
	if c.Previous == nil {
		err := "No endpoint to fetch previous locations from\n\n"
		fmt.Print(err)
		return errors.New(err)
	}

	locs, err := c.pokeapiClient.GetLocations(c.Previous)
	if err != nil {
		return err
	}

	c.Next = locs.Next
	c.Previous = locs.Previous
	for _, l := range locs.Results {
		fmt.Println(l.Name)
	}
	fmt.Print("\n")

	return nil
}

