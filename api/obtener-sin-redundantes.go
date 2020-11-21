package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/guillelpnz/TextAnalyzer/src/texto"
)

// Respuesta contains a text without duplicate words
type Respuesta struct {
	Contenido string `json:"texto"`
}

// Handler returns a webpage
func Handler(w http.ResponseWriter, r *http.Request) {
	var textoUnmarshall string

	defer r.Body.Close()
	var result map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal("Error desserializando json-> ", err)
		log.Printf("")
	}

	textoObj := texto.NewTextoRep(textoUnmarshall, "")
	contenidoSinR := ""

	for _, palabra := range textoObj.ObtenerSinRedundantes() {
		fmt.Fprintf(w, palabra)
		contenidoSinR += palabra
	}

	respSinSerializar := Respuesta{Contenido: contenidoSinR}

	respSerializada, _ := json.Marshal(respSinSerializar)

	w.Header().Add("Content-Type", "application/json")

	fmt.Fprintf(w, string(respSerializada))
}
