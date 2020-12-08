package micro

import (
	"testing"

	"github.com/guillelpnz/TextAnalyzer/src/texto"
)

func TestIntroducirTexto(t *testing.T) {
	array := NewTextos()

	cadenaPrueba := "Hola, esto es un texto de prueba"

	array.IntroducirTexto(cadenaPrueba, "Guillermo Lupiáñez")

	if array.Get(0).GetContenido() != cadenaPrueba ||
		array.Get(0).GetAutor() != "Guillermo Lupiáñez" {
		t.Error("Introducir texto mal implementado")
	}
}

func TestObtenerRedundantes(t *testing.T) {
	array := NewTextos()
	cadenaPrueba := "Hola, esto es un texto texto de prueba"
	cadenaEsperada := make([]string, 0)
	cadenaEsperada = append(cadenaEsperada, "texto")

	array.IntroducirTexto(cadenaPrueba, "Guillermo Lupiáñez")

	for i := 0; i < len(array.array); i++ {
		if cadenaEsperada[i] != array.ObtenerRedundantes(i)[i] {
			t.Error("Obtener Redundantes mal implementado")
		}
	}
}

func TestObtenerPersonas(t *testing.T) {
	array := NewTextos()

	cadena := "Hola, me llamo Guillermo. Ayer estuve visitando a mi amigo Aarón."

	array.IntroducirTexto(cadena, "Guillermo Lupiáñez")

	esperada := make([]string, 2)
	esperada[0] = "Guillermo"
	esperada[1] = "Aarón"

	i := 0
	for _, palabra := range array.ObtenerRedundantes(i) {
		if palabra != esperada[i] {
			t.Error("ObtenerPersonas mal implementado")
		}
		i++
	}
}

func TestObtenerSinRedundantes(t *testing.T) {
	array := NewTextos()
	cadena := "Mi nombre es Guillermo Lupiáñez. A veces practico pádel, otras veces practico Waterpolo"
	array.IntroducirTexto(cadena, "Guillermo Lupiáñez")

	esperada := "Mi nombre es Guillermo Lupiáñez. A veces practico pádel, otras Waterpolo"
	sliceEsperada := texto.StringToSlice(esperada)

	j := 0
	for i, palabra := range array.ObtenerSinRedundantes(j) {
		if palabra != sliceEsperada[i] {
			t.Error("Obtener redundantes mal implementado")
			break
		}
		j++
	}
}

func TestObtenerEstadisticas(t *testing.T) {
	array1 := NewTextos()

	cadena1 := "1 1 1 1 1 2 2 2 2 2"

	cadena2 := "1 1 1 1 2 2 2 3 3 3"

	array1.IntroducirTexto(cadena1, "Guillermo Lupiáñez")
	array1.IntroducirTexto(cadena2, "Guillermo Lupiáñez")

	dicc := make(map[string]float32, 2)

	dicc["1"] = 0.5
	dicc["2"] = 0.5

	est := array1.ObtenerEstadisticas(0)
	est2 := array1.ObtenerEstadisticas(1)

	if est["1"] != dicc["1"] || est["2"] != dicc["2"] {
		t.Error("Obtener estadisticas mal implementado")
	}

	dicc2 := make(map[string]float32, 3)

	dicc2["1"] = 0.4
	dicc2["2"] = 0.3
	dicc2["3"] = 0.3

	if est2["1"] != dicc2["1"] ||
		est2["2"] != dicc2["2"] ||
		est2["3"] != dicc2["3"] {
		t.Error("Obtener estadisticas mal implementado")
	}
}
