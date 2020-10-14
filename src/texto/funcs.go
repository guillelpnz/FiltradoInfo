package texto

func limpiar(texto string) string {
	signosPuntuacion := `.;:,"!¡?¿*-<>()|`
	var limpiada string

	for _, char := range texto {
		if !contains(signosPuntuacion, char) {
			limpiada = limpiada + string(char)
		}
	}

	return limpiada
}

func contains(haystack string, needle rune) bool {
	for _, char := range haystack {
		if needle == char {
			return true
		}
	}
	return false
}

func containsPalabra(haystack []string, needle string) bool {
	for _, palabra := range haystack {
		if needle == palabra {
			return true
		}
	}

	return false
}

func stringToSlice(texto string) []string {
	slice := make([]string, 1)
	var palabra string

	for _, char := range texto {
		if string(char) != " " {
			palabra = palabra + string(char)
		} else {
			slice = append(slice, palabra)
			palabra = ""
		}

	}

	return slice
}
