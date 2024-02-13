package pokeapi

import (
	"github.com/ViktorKharts/pokedex/internal/pokecache"
)

const (
	baseUrl = "https://pokeapi.co/api/v2/"
)

var cache = pokecache.NewCache(350) 

