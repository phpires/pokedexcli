package pokeapi

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
