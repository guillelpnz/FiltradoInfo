package texto

import "testing"

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
	var mapa map[string]int
	var autor string

	cadena := "Mi nombre es Guillermo Lupiáñez. A veces practico pádel, otras veces practico Waterpolo"
	esperada := "veces practico"
	sliceEsperada := StringToSlice(esperada)

	text := Texto{contenido: cadena, repeticiones: mapa, autor: autor}

	for i, palabra := range text.ObtenerSinRedundantes() {
		if palabra != sliceEsperada[i] {
			t.Error("Obtener redundantes mal implementado")
			break
		}
	}
}
