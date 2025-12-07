package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreaV2(url string) (LocationAreaResponseJson, error) {
	fmt.Printf("Making GET request from %v\n", url)
	if val, ok := c.cache.Get(url); ok {
		fmt.Println("Getting from cache")
		locationAreaResponseJson := LocationAreaResponseJson{}
		err := json.Unmarshal(val, &locationAreaResponseJson)
		if err != nil {
			return LocationAreaResponseJson{}, fmt.Errorf("error reading response from cache: %v", err)
		}
		return locationAreaResponseJson, nil
	}

	fmt.Println("Cache is empty, making pokeapi request")

	if len(url) == 0 {
		url = baseURL + "/location-area?offset=0&limit=20"
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponseJson{}, fmt.Errorf("error creating request: %v", err)
	}
	client := c.httpClient
	res, err := client.Do(req)
	if err != nil {
		return LocationAreaResponseJson{}, fmt.Errorf("error on request: %v", err)
	}
	defer res.Body.Close()

	statusCode := res.StatusCode
	if statusCode > 299 {
		return LocationAreaResponseJson{}, fmt.Errorf("status Code suggests error: %v", statusCode)
	}
	fmt.Printf("Request sucessful. Http status code: %v\n", statusCode)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponseJson{}, fmt.Errorf("error reading response body: %v", err)
	}

	var jsonResponse LocationAreaResponseJson
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return LocationAreaResponseJson{}, fmt.Errorf("error processing response body: %v", err)
	}

	c.cache.Add(url, body)

	return jsonResponse, nil
}
