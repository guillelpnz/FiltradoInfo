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
