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
	return t.Get(i).ObtenerRedundantes()
}

// ObtenerPersonas returns the allusions to people on a text (Spanish names)
func (t *Textos) ObtenerPersonas(i int) []string {
	return t.Get(i).ObtenerPersonas()
}

//ObtenerSinRedundantes returns the text in position i without repeated words
func (t *Textos) ObtenerSinRedundantes(i int) []string {
	return t.Get(i).ObtenerSinRedundantes()
}
