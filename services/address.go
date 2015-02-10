package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func Address2GPS(addr []rune, c *http.Client) (GPS, error) {
	url := "http://maps.googleapis.com/maps/api/geocode/json?sensor=false&address=" + string(addr)
	resp, err := c.Get(url)
	if err != nil {
		return GPS{}, errors.New("Get Nothing!!")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GPS{}, errors.New("Get Nothing!!")
	}
	var jsonobj MapAddressResult
	json.Unmarshal(body, &jsonobj)
	return jsonobj.Results[0].Geometry.Location, nil
}
