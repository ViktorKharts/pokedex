package pokeapi

import (
	"net/http"
	"io"
)

func get(c *Client, url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []byte{}, nil 
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, nil 
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, nil 
	}

	return data, nil 
}
