package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/guillelpnz/TextAnalyzer/src/texto"
)

// Respuesta contains a text without duplicate words
type Respuesta struct {
	Contenido string `json:"texto"`
}

// Peticion contains Unmarshaled request
// type Peticion struct {
// 	Contenido string `json:"texto"`
// }

// Handler returns a webpage
func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	textQuery := ""
	// var result Peticion

	//body, _ := ioutil.ReadAll(r.Body)

	// if err := json.Unmarshal(body, &result); err != nil {
	// 	log.Fatal("Error desserializando json-> ", err)
	// }

	switch r.Method {
	case "GET":
		for _, v := range r.URL.Query() {
			textQuery += v[0]
		}

		reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
		textQuery = reg.ReplaceAllString(textQuery, "")
		break
	}

	textoObj := texto.NewTextoRep(textQuery, "")
	contenidoSinR := ""

	for _, palabra := range textoObj.ObtenerSinRedundantes() {
		contenidoSinR += palabra + " "
	}

	respSinSerializar := Respuesta{Contenido: contenidoSinR}

	respSerializada, _ := json.Marshal(respSinSerializar)

	w.Header().Add("Content-Type", "application/json")

	fmt.Fprintf(w, "Texto con repetidos: %s", textQuery)
	fmt.Fprintf(w, "Texto sin repetidos: %s", string(respSerializada))
}
