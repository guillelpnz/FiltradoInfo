package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/guillelpnz/TextAnalyzer/src/texto"
)

// Respuesta contains a text without duplicate words
type Respuesta struct {
	Contenido string `json:"texto"`
}

// Handler processes a GET request and returns json data with no repeated words
func handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	textQuery := ""

	switch r.Method {
	case "GET":
		for _, v := range r.URL.Query() {
			textQuery += v[0]
		}

		reg, _ := regexp.Compile("[^a-zA-Z0-9\\s.,¡!¿?;<>]")
		textQuery = reg.ReplaceAllString(textQuery, "")
		break
	}

	textoObj := texto.NewTextoRep(textQuery, "")
	contenidoSinR := ""

	for _, palabra := range textoObj.ObtenerSinRedundantes() {
		contenidoSinR += palabra + " "
	}

	contenidoSinR = strings.TrimRight(contenidoSinR, " ")
	respSinSerializar := Respuesta{Contenido: contenidoSinR}

	respSerializada, _ := json.Marshal(respSinSerializar)

	w.Header().Add("Content-Type", "application/json")

	fmt.Fprintf(w, string(respSerializada))
}
