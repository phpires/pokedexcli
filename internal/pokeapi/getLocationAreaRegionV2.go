package pokeapi

import (
	"fmt"
)

func (c *Client) GetLocationAreaRegionV2(url string, regionName string) (LocationAreaRegionResponseJson, error) {
	if len(url) == 0 {
		url = baseURL + "/location-area"
	}
	if len(regionName) == 0 {
		return LocationAreaRegionResponseJson{}, fmt.Errorf("region name must be provided")
	}
	return LocationAreaRegionResponseJson{}, nil
}
