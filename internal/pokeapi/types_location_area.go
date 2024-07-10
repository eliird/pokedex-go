package pokeapi

type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`     //pointer to string because sometimes this can be null from the api response
	Previous *string `json:"previous"` //pointer to string because this can be null from the api response
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
