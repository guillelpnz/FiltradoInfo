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
