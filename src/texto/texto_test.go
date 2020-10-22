package texto

import (
	"testing"
)

func TestObtenerSinRedundantes(t *testing.T) {
	var mapa map[string]int
	var autor string

	cadena := "Mi nombre es Guillermo Lupiáñez. A veces practico pádel, otras veces practico Waterpolo"
	esperada := "Mi nombre es Guillermo Lupiáñez. A veces practico pádel, otras Waterpolo"
	sliceEsperada := StringToSlice(esperada)

	text := Texto{contenido: cadena, repeticiones: mapa, autor: autor}

	for i, palabra := range text.ObtenerSinRedundantes() {
		if palabra != sliceEsperada[i] {
			t.Error("Obtener redundantes mal implementado")
			break
		}
	}
}

func TestObtenerRedundantes(t *testing.T) {
	// var mapa map[string]int
	// var autor string

	cadena := "Mi nombre es Guillermo Lupiáñez. A veces practico pádel, otras veces practico Waterpolo"
	esperada := "veces practico"
	sliceEsperada := StringToSlice(esperada)

	text := NewTexto(cadena, nil, "")

	for i, palabra := range text.ObtenerRedundantes() {
		if palabra != sliceEsperada[i] {
			t.Error("Obtener redundantes mal implementado")
			break
		}
	}
}

// func TestObtenerPersonas(t *testing.T) {
// 	cadena := "Hola, me llamo Guillermo. Ayer estuve visitando a mi amigo Aarón."

// 	esperada := make([]string, 2)
// 	esperada[0] = "Guillermo"
// 	esperada[1] = "Aarón"
// 	text := NewTexto(cadena, nil, "")

// 	if text.ObtenerPersonas() != esperada {
// 		t.Error("ObtenerPersonas mal implementado")
// 	}
// }
