package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
	// var result Peticion

	body, _ := ioutil.ReadAll(r.Body)

	// if err := json.Unmarshal(body, &result); err != nil {
	// 	log.Fatal("Error desserializando json-> ", err)
	// }

	result := string(body)

	fmt.Fprintf(w, fmt.Sprintf("Result -> %s", result))

	textoObj := texto.NewTextoRep(result, "")
	contenidoSinR := ""

	for _, palabra := range textoObj.ObtenerSinRedundantes() {
		contenidoSinR += palabra + " "
	}

	respSinSerializar := Respuesta{Contenido: contenidoSinR}

	respSerializada, _ := json.Marshal(respSinSerializar)

	w.Header().Add("Content-Type", "application/json")

	fmt.Fprintf(w, "Texto sin repetidos -> "+string(respSerializada))
}
