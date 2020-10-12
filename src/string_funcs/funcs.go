package stringfuncs

func stringToSlice(texto string) {

	for i := 0; i < len(texto); i++ {
		if texto[i] == ` ` {

		}
	}
}

func contains(haystack string, needle rune) bool {
	for _, i := range haystack {
		if needle == i {
			return true
		}
	}
	return false
}
