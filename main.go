package main

import (
	"time"

	"github.com/ViktorKharts/pokedex/internal/pokeapi"
)

func main() {
	cfg := config{
		pokeapiClient:	pokeapi.NewClient(5 * time.Second),
		pokedex:	map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}
