//This package is used to send request to pokeapi
package pokeapi

import (
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"
 type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

type LocationAreasResp  struct {
	Count    int         `json:"count"`
	Next     *string      `json:"next"` //using *string because the value can be null - if null - nil else string value
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
} 