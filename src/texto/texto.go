package texto

import (
	"fmt"
)

type texto struct {
	contenido    string
	repeticiones map[string]int
	autor        string
}

func (t *texto) NewTexto(c string, r map[string]int, a string) *texto {
	t.contenido = c
	t.repeticiones = r
	t.autor = a

	return t
}
func (t *texto) GetContenido() string {
	return t.contenido
}

func (t *texto) GetRepeticiones() map[string]int {
	return t.repeticiones
}

func (t *texto) GetAutor() string {
	return t.autor
}

func (t *texto) SetContenido(contenido string) {
	t.contenido = contenido
}

func (t *texto) SetRepeticiones(repeticiones map[string]int) {
	t.repeticiones = repeticiones
}

func (t *texto) SetAutor(autor string) {
	t.autor = autor
}

func (t *texto) ObtenerRedundantes() []string {
	texto := Limpiar(t.contenido)
	slice := StringToSlice(texto)

	redundantes := make([]string, 0)

	size := len(slice)
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			if slice[i] == slice[j] && !ContainsPalabra(redundantes, slice[i]) {
				redundantes = append(redundantes, slice[i])
			}
		}
	}
	return redundantes
}

func (t *texto) IntroducirTexto(texto []string) {
	fmt.Println("Not implemented yet")
}

func (t *texto) GetReferencias() []string {
	s := make([]string, 1)

	fmt.Println("Not implemented yet")
	return s
}

func (t *texto) GetTextoPersona() []string {
	s := make([]string, 1)

	fmt.Println("Not implemented yet")
	return s
}
