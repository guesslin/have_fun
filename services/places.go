package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FindPlaces(location GPS, radius int32) []ResultObject {
	place_temp := "https://maps.googleapis.com/maps/api/place/search/json?radius=%d&sensor=false&key=%s&location=%f,%f"
	places := fmt.Sprintf(place_temp, radius, Key, location.Lat, location.Lon)
	resp, err := http.Get(places)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return nil
	}
	var placeresult PlacesResult
	json.Unmarshal(body, &placeresult)
	return placeresult.Results
}
