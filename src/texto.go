package texto

import "fmt"

type texto struct {
	palabras     []string
	repeticiones map[string]int
	autor        string
}

func (t *texto) getTexto() []string {
	return t.palabras
}

func (t *texto) getRepeticiones() map[string]int {
	return t.repeticiones
}

func (t *texto) getAutor() string {
	return t.autor
}

func (t *texto) setPalabras(palabras []string) {
	t.palabras = palabras
}

func (t *texto) setRepeticiones(repeticiones map[string]int) {
	t.repeticiones = repeticiones
}

func (t *texto) setAutor(autor string) {
	t.autor = autor
}

//Corresponde a HU1
func (t *texto) obtenerRedundantes() []string {
	s := make([]string, 1)
	fmt.Println("Not implemented yet")

	return s
}

//Corresponde a HU2
func (t *texto) introducirTexto(texto []string) {
	fmt.Println("Not implemented yet")
}

//Corresponde a HU3
func (t *texto) getReferencias() []string {
	s := make([]string, 1)

	fmt.Println("Not implemented yet")
	return s
}

//Corresponde a HU4
//el autor se guarda dentro de
func (t *texto) getTextoPersona() []string {
	s := make([]string, 1)

	fmt.Println("Not implemented yet")
	return s
}
