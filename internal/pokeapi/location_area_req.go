package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache

	data, ok := c.cache.Get(fullURL)
	if ok {
		// hit the cache
		fmt.Println("cache hit")
		locationAreasRespData := LocationAreasResp{}
		err := json.Unmarshal(data, &locationAreasRespData)
		if err != nil {
			return LocationAreasResp{}, err
		}

		return locationAreasRespData, nil
	}
	fmt.Println("no cache")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	// close response object
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasRespData := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasRespData)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreasRespData, nil
}
