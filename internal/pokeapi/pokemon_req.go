package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	// check the cache

	data, ok := c.cache.Get(fullURL)
	if ok {
		// hit the cache
		fmt.Println("cache hit")
		pokemonData := Pokemon{}
		err := json.Unmarshal(data, &pokemonData)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemonData, nil
	}
	fmt.Println("no cache")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	// close response object
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonData := Pokemon{}
	err = json.Unmarshal(data, &pokemonData)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)

	return pokemonData, nil
}
