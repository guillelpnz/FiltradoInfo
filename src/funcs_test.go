package texto

import "testing"

func TestContains(t *testing.T) {
	cadena := "Hola!"

	if !contains(cadena, '!') {
		t.Error("Contains no funciona adecuadamente")
	} else {
		t.Log("Contains funciona adecuadamente")
	}
}
func TestLimpiar(t *testing.T) {

}
