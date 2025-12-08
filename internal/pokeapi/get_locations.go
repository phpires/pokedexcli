package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocations(url string) (Location, error) {
	if val, ok := c.cache.Get(url); ok {
		locationAreaResponseJson := Location{}
		err := json.Unmarshal(val, &locationAreaResponseJson)
		if err != nil {
			return Location{}, fmt.Errorf("error reading response from cache: %v", err)
		}
		return locationAreaResponseJson, nil
	}

	if len(url) == 0 {
		url = baseURL + "/location-area?offset=0&limit=20"
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, fmt.Errorf("error creating request: %v", err)
	}
	client := c.httpClient
	res, err := client.Do(req)
	if err != nil {
		return Location{}, fmt.Errorf("error on request: %v", err)
	}
	defer res.Body.Close()

	statusCode := res.StatusCode
	if statusCode > 299 {
		return Location{}, fmt.Errorf("status Code suggests error: %v", statusCode)
	}
	fmt.Printf("Request sucessful. Http status code: %v\n", statusCode)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, fmt.Errorf("error reading response body: %v", err)
	}

	var jsonResponse Location
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return Location{}, fmt.Errorf("error processing response body: %v", err)
	}

	c.cache.Add(url, body)

	return jsonResponse, nil
}
