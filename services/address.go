package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Address2GPS(addr []rune) (float64, float64) {
	url := "http://maps.googleapis.com/maps/api/geocode/json?sensor=false&address=" + string(addr)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("GET NOTHING")
		return 0, 0
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("GET NOTHING")
		return 0, 0
	}
	var jsonobj MapAddressResult
	json.Unmarshal(body, &jsonobj)
	return jsonobj.Results[0].Geometry.Location.Lat, jsonobj.Results[0].Geometry.Location.Lon
}
