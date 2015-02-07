package greeting

import (
	"fmt"
	//	"github.com/guesslin/have_fun/services"
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
	// fmt.Fprint(w, "hello")
	err := placesTemplate.Execute(w, r.FormValue("address1"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var placesTemplate = template.Must(template.New("inplaces").Parse(placesTemplateHTML))
