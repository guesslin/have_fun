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
	Icon             string          `json:"icon"`
	ID               string          `json:"id"`
	Name             string          `json:"name"`
	PlaceID          string          `json:"place_id"`
	Reference        string          `json:"reference"`
	Scope            string          `json:"scope"`
	Vicinity         string          `json:"vicinity"`
}

type AddressObject struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type GeoObject struct {
	Location     GPS        `json:"location"`
	LocationType string     `json:"location_type"`
	ViewPort     ViewObject `json:"viewport"`
}

type GPS struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lng"`
}

type ViewObject struct {
	NoethEast GPS `json:"northeast"`
	SouthWest GPS `json:"southwest"`
}

type PlacesResult struct {
	HtmlAttributions []string       `json:"html_attributions"`
	NextPageToken    string         `json:"next_page_token"`
	Results          []ResultObject `json:"results"`
	Status           string         `json:"status"`
}
