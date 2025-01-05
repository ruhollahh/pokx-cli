package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RespPokemon struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ListPokemon(area string) (RespPokemon, error) {
	url := baseURL + fmt.Sprintf("/location-area/%s", area)

	data, exists := c.cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespPokemon{}, fmt.Errorf("newRequest: %w", err)
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespPokemon{}, fmt.Errorf("c.httpClient.Do: %w", err)
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespPokemon{}, fmt.Errorf("readAll: %w", err)
		}
		c.cache.Add(url, data)
	}

	var pokemons RespPokemon
	err := json.Unmarshal(data, &pokemons)
	if err != nil {
		return RespPokemon{}, fmt.Errorf("unmarshal: %w", err)
	}

	return pokemons, nil
}
