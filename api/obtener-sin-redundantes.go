package handler

import (
	"fmt"
	"net/http"

	"github.com/guillelpnz/TextAnalyzer/src/texto"
)

// Handler returns a webpage
func Handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "POST":
		r.ParseForm()

		textoForm := r.FormValue("texto")
		textoObj := texto.NewTextoRep(textoForm, "")

		for _, palabra := range textoObj.ObtenerSinRedundantes() {
			fmt.Fprintf(w, palabra)
		}

		break
	default:
		fmt.Fprintf(w, "Holaaaaa")
		http.ServeFile(w, r, "templates/index.html")
		break
	}
}
