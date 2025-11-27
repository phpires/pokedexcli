package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaResponseJson struct {
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Count    int       `json:"count"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"Name"`
	Url  string `json:"Url"`
}

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func GetLocationAreaV2(url string) (LocationAreaResponseJson, error) {

	if len(url) == 0 {
		url = baseURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponseJson{}, fmt.Errorf("error creating request: %v", err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return LocationAreaResponseJson{}, fmt.Errorf("error on request: %v", err)
	}
	defer res.Body.Close()

	statusCode := res.StatusCode
	if statusCode > 299 {
		return LocationAreaResponseJson{}, fmt.Errorf("status Code suggests error: %v", statusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponseJson{}, fmt.Errorf("error reading response body: %v", err)
	}

	var jsonResponse LocationAreaResponseJson
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return LocationAreaResponseJson{}, fmt.Errorf("error processing response body: %v", err)
	}

	return jsonResponse, nil
}
