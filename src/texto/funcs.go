package texto

//Limpiar returns a string without punctuation marks
func Limpiar(texto string) string {
	signosPuntuacion := `.;:,"!¡?¿*-<>()|`
	var limpiada string

	for _, char := range texto {
		if !Contains(signosPuntuacion, char) {
			limpiada = limpiada + string(char)
		}
	}

	return limpiada
}

// Contains returns true if needle is in haystack
func Contains(haystack string, needle rune) bool {
	for _, char := range haystack {
		if needle == char {
			return true
		}
	}
	return false
}

// ContainsPalabra returns true if needle is in haystack
func ContainsPalabra(haystack []string, needle string) bool {
	for _, palabra := range haystack {
		if needle == palabra {
			return true
		}
	}

	return false
}

//StringToSlice returns a slice from a string
func StringToSlice(texto string) []string {
	var slice []string
	var palabra string
	size := len(texto)

	for i, char := range texto {

		if string(char) != " " && string(char) != "\n" {
			palabra = palabra + string(char)
		} else {
			slice = append(slice, palabra)
			palabra = ""
		}

		if i == size-1 {
			slice = append(slice, palabra)
		}
	}

	return slice
}

//Check panics in case of error
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
