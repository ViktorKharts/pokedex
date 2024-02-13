package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + pokemonName

	p := Pokemon{}
	if cachedP, ok := cache.Get(url); ok {
		err := json.Unmarshal(cachedP, &p)
		if err == nil {
			return p, nil
		}
	}

	data, err := get(c, url)
	if err != nil {
		return p, err
	}
	cache.Add(url, data)

	if err = json.Unmarshal(data, &p); err != nil {
		return p, err
	}

	return p, nil
}

