package texto

import "testing"

func TestContains(t *testing.T) {
	var cadena string
	cadena = "Hola!"

	if !Contains(cadena, '!') {
		t.Error("Contains no funciona adecuadamente")
	}
}
