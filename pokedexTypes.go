package main

import (
	"github.com/ViktorKharts/pokedex/internal/pokeapi"
)

type commands map[string]cliCommand

type cliCommand struct {
	name		string
	description 	string
	callback	func(*config, []string) error
}

type config struct {
	pokeapiClient	pokeapi.Client
	pokedex		map[string]pokeapi.Pokemon
	Next		*string 
	Previous	*string
}

