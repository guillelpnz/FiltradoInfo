package texto

import "testing"

func TestContains(t *testing.T) {
	var cadena string
	cadena = "Hola!"

	if !contains(cadena, '!') {
		t.Error("Contains no funciona adecuadamente")
	}
}
func TestLimpiar(t *testing.T) {
	var cadena, correcta string

	cadena = "¡¡Hola!! Mi nombre es Guille."
	correcta = "Hola Mi nombre es Guille"

	cadena = limpiar(cadena)

	if cadena != correcta {
		t.Error("Limpiar no funciona adecuadamente")
	}
}

func TestContainsPalabra(t *testing.T) {
	var slice []string

	slice = make([]string, 4)
	slice[0] = "Hola"
	slice[1] = "me"
	slice[2] = "llamo"
	slice[3] = "Guillermo"

	if !containsPalabra(slice, "me") {
		t.Error("ContainsPalabra no funciona adecuadamente")
	}
}
