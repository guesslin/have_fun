package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Address2GPS(addr []rune, c *http.Client) GPS {
	url := "http://maps.googleapis.com/maps/api/geocode/json?sensor=false&address=" + string(addr)
	resp, err := c.Get(url)
	if err != nil {
		fmt.Println("GET NOTHING")
		return GPS{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("GET NOTHING")
		return GPS{}
	}
	var jsonobj MapAddressResult
	json.Unmarshal(body, &jsonobj)
	return jsonobj.Results[0].Geometry.Location
}
