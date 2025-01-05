package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RespAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListAreas(pageURL *string) (RespAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, exists := c.cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespAreas{}, fmt.Errorf("http.NewRequest: %w", err)
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespAreas{}, fmt.Errorf("c.httpClient.Do: %w", err)
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespAreas{}, fmt.Errorf("io.ReadAll: %w", err)
		}
		c.cache.Add(url, data)
	}

	var areas RespAreas
	err := json.Unmarshal(data, &areas)
	if err != nil {
		return RespAreas{}, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return areas, nil
}
