package texto

import (
	"testing"
)

func TestObtenerSinRedundantes(t *testing.T) {
	var mapa map[string]int
	var autor string

	cadena := "Mi nombre es Guillermo Lupiáñez. A veces practico pádel, otras veces practico Waterpolo"
	esperada := "Mi nombre es Guillermo Lupiáñez. A veces practico pádel, otras Waterpolo"
	sliceEsperada := StringToSlice(esperada)

	text := Texto{contenido: cadena, repeticiones: mapa, autor: autor}

	for i, palabra := range text.ObtenerSinRedundantes() {
		if palabra != sliceEsperada[i] {
			t.Error("Obtener redundantes mal implementado")
			break
		}
	}
}

func TestObtenerRedundantes(t *testing.T) {
	// var mapa map[string]int
	// var autor string

	cadena := "Mi nombre es Guillermo Lupiáñez. A veces practico pádel, otras veces practico Waterpolo"
	esperada := "veces practico"
	sliceEsperada := StringToSlice(esperada)

	text := NewTexto(cadena, nil, "")

	for i, palabra := range text.ObtenerRedundantes() {
		if palabra != sliceEsperada[i] {
			t.Error("Obtener redundantes mal implementado")
			break
		}
	}
}

// func TestObtenerPersonas(t *testing.T) {

// 	text := NewTexto("", nil, "")
// 	t.Log(text.ObtenerPersonas())
// }

// func TestObtenerPersonas(t *testing.T) {
// 	cadena := "Hola, me llamo Guillermo. Ayer estuve visitando a mi amigo Aarón."

// 	esperada := make([]string, 2)
// 	esperada[0] = "Guillermo"
// 	esperada[1] = "Aarón"
// 	text := NewTexto(cadena, nil, "")

// 	for i, palabra := range text.ObtenerPersonas() {
// 		if palabra != esperada[i] {
// 			t.Error("Obtener personas mal implementado")
// 			break
// 		}
// 	}
// }

func TestNewTextoRep(t *testing.T) {
	text := NewTextoRep("hola que tal!!, hola que tal!!", "")

	mapa := text.GetRepeticiones()

	if mapa["hola"] != 2 || mapa["que"] != 2 || mapa["tal"] != 2 || mapa["!"] != 0 {
		t.Error("Constructor repes mal implementado")
	}

}

func TestObtenerEstadisticas(t *testing.T) {
	textoTest := NewTextoRep("1 1 1 1 1 2 2 2 2 2", "")
	textoTest2 := NewTextoRep("1 1 1 1 2 2 2 3 3 3", "")

	dicc := make(map[string]float32, 2)

	dicc["1"] = 0.5
	dicc["2"] = 0.5

	//est := make(map[string]float32, 2)
	est := textoTest.ObtenerEstadisticas()

	if est["1"] != dicc["1"] || est["2"] != dicc["2"] {
		t.Error("Obtener estadisticas mal implementado")
	}

	dicc2 := make(map[string]float32, 3)

	dicc2["1"] = 0.4
	dicc2["2"] = 0.3
	dicc2["3"] = 0.3

	//est2 := make(map[string]float32, 3)
	est2 := textoTest2.ObtenerEstadisticas()

	if est2["1"] != dicc2["1"] ||
		est2["2"] != dicc2["2"] ||
		est2["3"] != dicc2["3"] {
		t.Error("Obtener estadisticas mal implementado")
	}

}

func TestNewTexto(t *testing.T) {
	dicc := make(map[string]int, 2)
	dicc["1"] = 1

	text := NewTexto("hola que tal!!", dicc, "Bilal")

	if text.GetContenido() != "hola que tal!!" || text.repeticiones["1"] != 1 || text.autor != "Bilal" {
		t.Error("Constructor simple mal implementado")
	}

}

func TestGetContenido(t *testing.T) {
	dicc := make(map[string]int, 2)
	dicc["1"] = 1
	text := NewTexto("hola que tal!!", dicc, "Bilal")
	text.SetContenido("Estoy bien")

	if text.GetContenido() != "Estoy bien" {
		t.Error("GetContenido() mal implementado")
	}
}

func TestGetRepeticiones(t *testing.T) {
	dicc := make(map[string]int, 2)
	dicc["1"] = 1
	otroDicc := make(map[string]int, 2)
	otroDicc["1"] = 4
	otroDicc["2"] = 5
	text := NewTexto("hola que tal!!", dicc, "Bilal")
	text.SetRepeticiones(otroDicc)

	if text.GetRepeticiones()["1"] != 4 || text.GetRepeticiones()["2"] != 5 {
		t.Error("GetRepeticiones() mal implementado")
	}
}

func TestGetAutor(t *testing.T) {
	dicc := make(map[string]int, 2)
	dicc["1"] = 1
	text := NewTexto("hola que tal!!", dicc, "Bilal")
	text.SetAutor("Mohamed")

	if text.GetAutor() != "Mohamed" {
		t.Error("Get() mal implementado")
	}
}

func TestObtenerNumPalabras(t *testing.T) {
	text := NewTextoRep("hola que tal!!", "")
	num := text.ObtenerNumPalabras()

	if num != 3 {
		t.Error("ObtenerNumPalabras() mal implementado")
	}
}
