package pokeapi

import (
	"net/http"
	"time"

	"github.com/phpires/pokedexcli/internal/pokecache"
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

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, interval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
