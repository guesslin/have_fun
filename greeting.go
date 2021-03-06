package greeting

import (
	"appengine"
	"appengine/urlfetch"
	"github.com/guesslin/services"
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/inplaces", inplaces)
}

func inplaces(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	point1, err := services.Address2GPS([]rune(r.FormValue("address1")), client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	point2, err := services.Address2GPS([]rune(r.FormValue("address2")), client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	midpoint := services.LookingForMidpoint([]services.GPS{point1, point2})
	results := services.FindPlaces(midpoint, 100, client)
	err = placesTemplate.Execute(w, results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var placesTemplate = template.Must(template.New("inplaces").Parse(placesTemplateHTML))
