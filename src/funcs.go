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
