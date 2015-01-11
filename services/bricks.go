package services

type MapAddressResult struct {
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
