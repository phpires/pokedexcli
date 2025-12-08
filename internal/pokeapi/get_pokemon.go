package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (PokemonInfo, error) {
	url := baseURL + "/pokemon/" + name

	if len(name) == 0 {
		return PokemonInfo{}, fmt.Errorf("pokemon name must be provided")
	}

	if val, ok := c.cache.Get(url); ok {
		cacheValue := PokemonInfo{}
		err := json.Unmarshal(val, &cacheValue)
		if err != nil {
			return PokemonInfo{}, fmt.Errorf("error reading from cache: %v", err)
		}

		return cacheValue, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, fmt.Errorf("error building request: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()

	statusCode := res.StatusCode

	if statusCode == 404 {
		return PokemonInfo{}, fmt.Errorf("pokemon not found: %v", name)
	}

	if statusCode > 299 {
		return PokemonInfo{}, fmt.Errorf("status Code suggests error: %v", statusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonInfo{}, fmt.Errorf("error reading response body: %v", err)
	}

	var jsonResponse PokemonInfo
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return PokemonInfo{}, fmt.Errorf("error decoding body response: %v", err)
	}

	c.cache.Add(url, body)

	return jsonResponse, nil
}
