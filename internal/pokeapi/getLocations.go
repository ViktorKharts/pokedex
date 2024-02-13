package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetLocations(pageUrl *string) (Locations, error) {
	url := baseUrl + "/location-area" 
	if pageUrl != nil {
		url = *pageUrl
	}

	locs := Locations{}
	if cachedLocs, ok := cache.Get(url); ok {
		err := json.Unmarshal(cachedLocs, &locs)
		if err == nil {
			return locs, nil
		}
	}

	data, err := get(c, url)
	if err != nil {
		return locs, err
	}
	cache.Add(url, data)

	if err := json.Unmarshal(data, &locs); err != nil {
		return locs, err
	}

	return locs, nil
}
