package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RespPokemon struct {
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"base_experience"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func (c *Client) GetPokemon(name string) (RespPokemon, error) {
	url := baseURL + fmt.Sprintf("/pokemon/%s", name)
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

	var pokemon RespPokemon
	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return RespPokemon{}, fmt.Errorf("unmarshal: %w", err)
	}

	return pokemon, nil
}
