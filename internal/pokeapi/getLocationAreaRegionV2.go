package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreaRegionV2(regionName string) (LocationAreaRegionResponseJson, error) {

	reqUrl := fmt.Sprintf("%s/location-area/%s", baseURL, regionName)

	if len(regionName) == 0 {
		return LocationAreaRegionResponseJson{}, fmt.Errorf("region name must be provided")
	}
	fmt.Printf("The url to request is %s\n", reqUrl)

	if val, ok := c.cache.Get(reqUrl); ok {
		fmt.Println("Found cached value.")
		cacheValue := LocationAreaRegionResponseJson{}
		err := json.Unmarshal(val, &cacheValue)
		if err != nil {
			return LocationAreaRegionResponseJson{}, fmt.Errorf("error reading from cache: %v", err)
		}
		return cacheValue, nil
	}

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return LocationAreaRegionResponseJson{}, fmt.Errorf("error building request: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaRegionResponseJson{}, fmt.Errorf("error executing request: %v", err)
	}
	defer res.Body.Close()

	statusCode := res.StatusCode
	if statusCode > 299 {
		return LocationAreaRegionResponseJson{}, fmt.Errorf("status Code suggests error: %v", statusCode)
	}
	fmt.Printf("Request sucessful. Http status code: %v\n", statusCode)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaRegionResponseJson{}, fmt.Errorf("error reading body response: %v", statusCode)
	}

	var jsonResponse LocationAreaRegionResponseJson
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return LocationAreaRegionResponseJson{}, fmt.Errorf("error decoding body response: %v", statusCode)
	}
	c.cache.Add(reqUrl, body)
	return jsonResponse, nil
}
