package texto

import "testing"

func TestContains(t *testing.T) {
	var cadena string
	cadena = "Hola!"

	if !Contains(cadena, '!') {
		t.Error("Contains no funciona adecuadamente")
	}
}

func TestLimpiar(t *testing.T) {
	var cadena, correcta string

	cadena = "¡¡Hola!! Mi nombre es Guille.."
	correcta = "Hola Mi nombre es Guille"

	cadena = Limpiar(cadena)

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

	if !ContainsPalabra(slice, "Guillermo") {
		t.Error("ContainsPalabra no funciona adecuadamente")
	}
}

func TestStringToSlice(t *testing.T) {
	var cadena1 string
	cadena1 = "Hola esto es una cadena de prueba"

	correcto := make([]string, 7)
	correcto[0] = "Hola"
	correcto[1] = "esto"
	correcto[2] = "es"
	correcto[3] = "una"
	correcto[4] = "cadena"
	correcto[5] = "de"
	correcto[6] = "prueba"

	for i, palabra := range StringToSlice(cadena1) {
		if palabra != correcto[i] {
			t.Error("StringToSlice no funciona adecuadamente")
			break
		}
	}

}
