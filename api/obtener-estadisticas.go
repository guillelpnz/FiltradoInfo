//Package handler contains an HTTP Cloud Function.
package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/guillelpnz/TextAnalyzer/src/texto"
)

// Response stores a telegram response
type Response struct {
	Msg    string `json:"text"`
	ChatID int64  `json:"chat_id"`
	Method string `json:"method"`
}

//Peticion stores the poem content
type Peticion struct {
	Contenido string `json:"texto"`
}

//Poema stores a poem content
type Poema struct {
	Contenido string `json:"texto"`
}

var poema = Poema{
	Contenido: "Mis pasos en esta calle Resuenan En otra calle Donde Oigo mis pasos Pasar en esta calle Donde Sólo es real la niebla.",
}

// Handler returns the statistics of a famous poem in Spanish
func Handler(w http.ResponseWriter, r *http.Request) {
	var poema Poema
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var update tgbotapi.Update
	if err := json.Unmarshal(body, &update); err != nil {
		log.Fatal("Error en el update →", err)
	}
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	if update.Message.IsCommand() {
		text := ""
		switch update.Message.Command() {
		case "estadisticas":
			fmt.Printf(poema.Contenido)
			textoObj := texto.NewTextoRep(poema.Contenido, "Antonio Paz")
			text = "Estadísticas del poema Aquí -> "

			for key, entrada := range textoObj.ObtenerEstadisticas() {
				text += key + " : " + fmt.Sprintf("%f", entrada) + ";"
			}
		default:
			text = "Usa /estadisticas para obtener estadisticas sobre el poema: Aquí de Antonio Paz"
		}
		data := Response{Msg: text,
			Method: "sendMessage",
			ChatID: update.Message.Chat.ID}

		msg, _ := json.Marshal(data)
		log.Printf("Response %s", string(msg))
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, string(msg))
	}
}
