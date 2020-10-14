package main

import (
	"fmt"

	"github.com/guillelpnz/TextAnalyzer/src/texto"
)

func main() {
	cadena := "Mi nombre es Guillermo Lupiáñez. A veces practico pádel, otras veces practico Waterpolo"
	esperada := make([]string, 2)
	esperada[0] = "veces"
	esperada[1] = "practico"

	text := texto.NewTexto(cadena, nil, "")

	fmt.Println(text.ObtenerSinRedundantes())
	// for i, palabra := range text.ObtenerRedundantes() {
	// 	if palabra != esperada[i] {
	// 		fmt.Println("Hola")
	// 	}
	// }

}
