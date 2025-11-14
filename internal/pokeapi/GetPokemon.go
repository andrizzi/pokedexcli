package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (RespDetailedPokemons, error) {
	url := baseURL + "/pokemon/" + name

	// check cache first
	if cachedData, found := c.cache.Get(url); found {
		var locationResp RespDetailedPokemons
		err := json.Unmarshal(cachedData, &locationResp)
		if err == nil {
			return locationResp, nil
		}
		// if unmarshal fails, proceed to fetch from API
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDetailedPokemons{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDetailedPokemons{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDetailedPokemons{}, err
	}

	var locationsResp RespDetailedPokemons
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespDetailedPokemons{}, err
	}

	c.cache.Add(url, dat)

	return locationsResp, nil
}
