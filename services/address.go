package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JsonResult struct {
	Results []ResultObject `json:"results"`
	Status  string         `json:"status"`
}

type ResultObject struct {
	AddressComponent []AddressObject `json:"address_components"`
	FormattedAddress string          `json:"formatted_address"`
	Geometry         GeoObject       `json:"geometry"`
	PartialMatch     bool            `json:"partial_match"`
	Types            []string        `json:"types"`
}

type AddressObject struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type GeoObject struct {
	Location     GPSObject  `json:"location"`
	LocationType string     `json:"location_type"`
	ViewPort     ViewObject `json:"viewport"`
}

type GPSObject struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lng"`
}

type ViewObject struct {
	NoethEast GPSObject `json:"northeast"`
	SouthWest GPSObject `json:"southwest"`
}

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
	var jsonobj JsonResult
	json.Unmarshal(body, &jsonobj)
	return jsonobj.Results[0].Geometry.Location.Lat, jsonobj.Results[0].Geometry.Location.Lon
}
