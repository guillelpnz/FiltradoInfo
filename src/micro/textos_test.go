package micro

import (
	"testing"
)

func TestIntroducirTexto(t *testing.T) {
	array := NewTextos()

	array.IntroducirTexto("Hola, esto es un texto de prueba", "Guillermo Lupiáñez")
}
