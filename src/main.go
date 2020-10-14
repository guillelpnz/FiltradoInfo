package main

import "github.com/guillelpnz/TextAnalyzer/src/texto"
import "fmt"

func main() {
	slice := texto.StringToSlice("Holaaaa me llamo Guille")

	for _,palabra := range slice {
		fmt.Println(palabra)
	}
}
