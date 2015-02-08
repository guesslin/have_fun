package greeting

import (
	"fmt"
	"github.com/guesslin/have_fun/services"
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/inplaces", inplaces)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, AddresesForm)
}

func inplaces(w http.ResponseWriter, r *http.Request) {
	point1 := services.Address2GPS([]rune(r.FormValue("address1")))
	point2 := services.Address2GPS([]rune(r.FormValue("address2")))
	fmt.Fprint(w, point2.Lat)
	midpoint := services.LookingForMidpoint(point1, point2)
	results := services.FindPlaces(midpoint, 100)
	err := placesTemplate.Execute(w, results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var placesTemplate = template.Must(template.New("inplaces").Parse(placesTemplateHTML))
