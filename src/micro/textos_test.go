package micro

import (
	"testing"
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
