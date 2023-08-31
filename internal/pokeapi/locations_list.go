package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(PageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if PageURL != nil {
		url = *PageURL
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return RespLocations{}, err
	}

	locationsResp := RespLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespLocations{}, err
	}
	return locationsResp, nil
}
