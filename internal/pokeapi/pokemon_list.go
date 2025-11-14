package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(name string) (RespDetailedLocationArea, error) {
	url := baseURL + "/location-area/" + name

	// check cache first
	if cachedData, found := c.cache.Get(url); found {
		var locationResp RespDetailedLocationArea
		err := json.Unmarshal(cachedData, &locationResp)
		if err == nil {
			return locationResp, nil
		}
		// if unmarshal fails, proceed to fetch from API
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDetailedLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDetailedLocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDetailedLocationArea{}, err
	}

	var locationsResp RespDetailedLocationArea
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespDetailedLocationArea{}, err
	}

	c.cache.Add(url, dat)

	return locationsResp, nil
}
