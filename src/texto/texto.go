package texto

import (
	"fmt"
	"io/ioutil"
)

//Texto struct to store information about a text
type Texto struct {
	contenido    string
	repeticiones map[string]int
	autor        string
}

//NewTexto returns a new texto object
func NewTexto(c string, r map[string]int, a string) *Texto {
	t := new(Texto)
	t.contenido = c
	t.repeticiones = r
	t.autor = a

	return t
}

//NewTextoRep returns a new texto object with @repeticiones correctly initialized
func NewTextoRep(c string, a string) *Texto {
	t := new(Texto)
	t.contenido = c
	t.autor = a

	cadenaContenido := Limpiar(t.contenido)
	sliceContenido := StringToSlice(cadenaContenido)

	t.repeticiones = make(map[string]int)
	for _, p := range sliceContenido {
		t.repeticiones[p]++
	}

	return t
}

//GetContenido returns Texto.contenido
func (t *Texto) GetContenido() string {
	return t.contenido
}

//GetRepeticiones returns Texto.repeticiones
func (t *Texto) GetRepeticiones() map[string]int {
	return t.repeticiones
}

//GetAutor return the author of a Texto object
func (t *Texto) GetAutor() string {
	return t.autor
}

//SetContenido sets the attribute contenido
func (t *Texto) SetContenido(contenido string) {
	t.contenido = contenido
}

//SetRepeticiones sets the attribute repeticiones
func (t *Texto) SetRepeticiones(repeticiones map[string]int) {
	t.repeticiones = repeticiones
}

//SetAutor sets the attribute autor
func (t *Texto) SetAutor(autor string) {
	t.autor = autor
}

//ObtenerSinRedundantes returns a text without repeated words
func (t *Texto) ObtenerSinRedundantes() []string {
	signosPuntuacion := `.;:,"!¡?¿*-<>()|`

	slice := StringToSlice(t.contenido)

	sinRedundantes := make([]string, 0)

	size := len(slice)
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			if !ContainsPalabra(sinRedundantes, slice[i]) {
				for _, char := range signosPuntuacion {
					if !Contains(slice[i], char) {
						sinRedundantes = append(sinRedundantes, slice[i])
						break
					}
				}

			}
		}
	}
	return sinRedundantes
}

//ObtenerRedundantes returns the repeated words on a text
func (t *Texto) ObtenerRedundantes() []string {
	limpio := Limpiar(t.contenido)

	slice := StringToSlice(limpio)
	var redundantes []string

	for i, palabra1 := range slice {
		for j := i + 1; j < len(slice); j++ {
			if palabra1 == slice[j] && !ContainsPalabra(redundantes, palabra1) {
				redundantes = append(redundantes, palabra1)
			}
		}
	}

	return redundantes
}

//IntroducirTexto lets the user store text in the db
func (t *Texto) IntroducirTexto(texto []string) {
	fmt.Println("Not implemented yet")
}

//GetTextoPersona returns the texts a person has written
func (t *Texto) GetTextoPersona() []string {
	s := make([]string, 1)

	fmt.Println("Not implemented yet")
	return s
}

//ObtenerPersonas returns the allusions to people in a text
func (t *Texto) ObtenerPersonas() []string {
	var personas []string
	diccionario, err := ioutil.ReadFile("./diccionarios/diccionario_propios.txt")
	Check(err)

	sliceNombres := StringToSlice(string(diccionario))
	contenido := Limpiar(t.contenido)
	sliceContenido := StringToSlice(contenido)

	for i := 0; i < len(sliceContenido); i++ {
		for j := i; j < len(sliceNombres); j++ {
			if sliceContenido[i] == sliceNombres[j] {
				personas = append(personas, sliceContenido[i])
			}
		}
	}

	return personas
}

//ObtenerEstadisticas returns the percentage of use of each word in a text
func (t *Texto) ObtenerEstadisticas() map[string]float32 {
	var numPalabras int
	numPalabras = 0
	estadisticas := make(map[string]float32)
	var palabras []string
	for _, num := range t.repeticiones {
		numPalabras += num
	}

	for palabra, num := range t.repeticiones {
		repetida := false
		for _, p := range palabras {
			if palabra == p {
				repetida = true
				break
			}
		}

		if !repetida {
			estadisticas[palabra] = float32(num) / float32(numPalabras)
			palabras = append(palabras, palabra)
		}

	}

	return estadisticas
}

// ObtenerNumPalabras returns the number of words in a text
func (t *Texto) ObtenerNumPalabras() int {
	return len(StringToSlice(t.contenido))
}
