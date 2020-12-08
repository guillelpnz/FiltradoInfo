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
