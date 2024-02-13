package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetPokemons(location *string) (AreaPokemons, error) {
	url := baseUrl + "/location-area/" + *location 

	ap := AreaPokemons{}
	if cachedAp, ok := cache.Get(url); ok {
		err := json.Unmarshal(cachedAp, &ap)
		if err == nil {
			return ap, nil
		}
	}

	data, err := get(c, url)
	if err != nil {
		return ap, err
	}
	cache.Add(url, data)

	if err := json.Unmarshal(data, &ap); err != nil {
		return ap, err
	}

	return ap, nil
}

