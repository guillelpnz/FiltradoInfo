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

//IntroducirTexto adds a new text to Textos.array
func (t *Textos) IntroducirTexto(contenido string, autor string) {
	nuevo := texto.NewTextoRep(contenido, autor)

	t.array = append(t.array, *nuevo)
}
