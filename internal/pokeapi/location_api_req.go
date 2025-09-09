
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
func (c *Client) LocationAreas() (LocationAreasResp, error) {


	url := "/location-area"
	full_url := baseURL + url

	 req, err := http.NewRequest("GET", full_url, nil)
	 if err != nil {
		return LocationAreasResp{}, err
	 }
	 
	 resp, err := c.httpClient.Do(req) 

	 defer resp.Body.Close()

	 if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code %v",resp.StatusCode)

	 }

	 data, err := io.ReadAll(resp.Body)

	 if err != nil {
		return LocationAreasResp{}, err
	 }

	 // this function loads the data (json) into struct

	 locationAreasResp := LocationAreasResp{}
	 err = json.Unmarshal(data, &locationAreasResp)

	 if err != nil {
		return LocationAreasResp{}, err
	 } 

	 return locationAreasResp, nil
	  



}