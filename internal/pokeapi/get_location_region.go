package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationRegion(regionName string) (LocationRegion, error) {

	reqUrl := fmt.Sprintf("%s/location-area/%s", baseURL, regionName)

	if len(regionName) == 0 {
		return LocationRegion{}, fmt.Errorf("region name must be provided")
	}

	if val, ok := c.cache.Get(reqUrl); ok {
		fmt.Println("Found cached value.")
		cacheValue := LocationRegion{}
		err := json.Unmarshal(val, &cacheValue)
		if err != nil {
			return LocationRegion{}, fmt.Errorf("error reading from cache: %v", err)
		}
		return cacheValue, nil
	}

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return LocationRegion{}, fmt.Errorf("error building request: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationRegion{}, fmt.Errorf("error executing request: %v", err)
	}
	defer res.Body.Close()

	statusCode := res.StatusCode
	if statusCode > 299 {
		return LocationRegion{}, fmt.Errorf("status Code suggests error: %v", statusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationRegion{}, fmt.Errorf("error reading body response: %v", statusCode)
	}

	var jsonResponse LocationRegion
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return LocationRegion{}, fmt.Errorf("error decoding body response: %v", statusCode)
	}
	c.cache.Add(reqUrl, body)
	return jsonResponse, nil
}
