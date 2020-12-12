package micro

import "github.com/guillelpnz/TextAnalyzer/src/texto"

// Textos contains a slice of Texto objects
type Textos struct {
	array []texto.Texto
}

// NewTextos is the default constructor of class Textos
func NewTextos() *Textos {
	t := new(Textos)

	return t
}

// Get returns the text at i position of the array
func (t *Textos) Get(i int) *texto.Texto {
	return &t.array[i]
}

//IntroducirTexto adds a new text to Textos.array
func (t *Textos) IntroducirTexto(contenido string, autor string) {
	nuevo := texto.NewTextoRep(contenido, autor)

	t.array = append(t.array, *nuevo)
}

// ObtenerRedundantes returns the repeated words in the text on i position
func (t *Textos) ObtenerRedundantes(i int) []string {
	if i >= 0 && i < len(t.array) {
		return t.Get(i).ObtenerRedundantes()
	} else {
		array := make([]string, 0)
		array = texto.StringToSlice("index out of array")
		return array
	}
}

// ObtenerPersonas returns the allusions to people on a text (Spanish names)
func (t *Textos) ObtenerPersonas(i int) []string {
	if i >= 0 && i < len(t.array) {
		return t.Get(i).ObtenerPersonas()
	} else {
		array := make([]string, 0)
		array = texto.StringToSlice("index out of array")
		return array
	}
}

//ObtenerSinRedundantes returns the text in position i without repeated words
func (t *Textos) ObtenerSinRedundantes(i int) []string {
	if i >= 0 && i < len(t.array) {
		return t.Get(i).ObtenerSinRedundantes()
	} else {
		array := make([]string, 0)
		array = texto.StringToSlice("index out of array")
		return array
	}
}

//ObtenerEstadisticas returns the text in position i without repeated words
func (t *Textos) ObtenerEstadisticas(i int) map[string]float32 {
	if i >= 0 && i < len(t.array) {
		return t.Get(i).ObtenerEstadisticas()
	} else {
		array := make(map[string]float32, 0)
		array["index out of array"] = 1
		return array
	}
}
