package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIntroducirTexto(t *testing.T) {
	r := setupServer()
	w := httptest.NewRecorder()

	param := strings.NewReader("autor=Guillermo Lupiáñez Tapia&contenido=mi nombre es Guillermo Guillermo y esto es una prueba")
	req, err := http.NewRequest("POST", "/texto", param)

	if err != nil {
		t.Fatalf("No se esperaba ningún error, se obtuvo %v", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	//resp, err := http.Get("http://localhost:8080/obtener-redundantes?posicion=0")

	if w.Code != 200 {
		t.Fatalf("Se esperaba código=200, se obtuvo %d", w.Code)
	}

	want := "{\"autor\":\"Guillermo Lupiáñez Tapia\",\"contenido\":\"mi nombre es Guillermo Guillermo y esto es una prueba\",\"mensaje\":\"Texto añadido con éxito\"}"
	if w.Body.String() != want {
		t.Fatalf("Error introduciendo texto, cadena esperada: %v. Cadena obtenida %v", want, w.Body.String())
	}
}

func TestObtenerRedundantes(t *testing.T) {
	r := setupServer()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/obtener-redundantes?posicion=0", nil)

	r.ServeHTTP(w, req)
	//resp, err := http.Get("http://localhost:8080/obtener-redundantes?posicion=0")

	if err != nil {
		t.Fatalf("No se esperaba ningún error, se obtuvo %v", err)
	}

	if w.Code != 200 {
		t.Fatalf("Se esperaba código=200, se obtuvo %d", w.Code)
	}

	want := "{\"mensaje\":\"Texto analizado con éxito\",\"palabras redundantes\":[\"index\",\"out\",\"of\",\"array\"]}"

	if want != w.Body.String() {
		t.Fatalf("Error obteniendo redundantes, cadena esperada: %v. Cadena obtenida %v", want, w.Body.String())
	}
}

func TestObtenerSinRedundantes(t *testing.T) {
	r := setupServer()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/obtener-sin-redundantes?posicion=0", nil)

	r.ServeHTTP(w, req)
	//resp, err := http.Get("http://localhost:8080/obtener-redundantes?posicion=0")

	if err != nil {
		t.Fatalf("No se esperaba ningún error, se obtuvo %v", err)
	}

	if w.Code != 200 {
		t.Fatalf("Se esperaba código=200, se obtuvo %d", w.Code)
	}

	want := "{\"Texto sin redundantes\":[\"index\",\"out\",\"of\",\"array\"],\"mensaje\":\"Texto analizado con éxito\"}"

	if want != w.Body.String() {
		t.Fatalf("Error obteniendo sin redundantes, cadena esperada: %v. Cadena obtenida %v", want, w.Body.String())
	}
}

func TestObtenerPersonas(t *testing.T) {
	r := setupServer()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/obtener-personas?posicion=0", nil)

	r.ServeHTTP(w, req)
	//resp, err := http.Get("http://localhost:8080/obtener-redundantes?posicion=0")

	if err != nil {
		t.Fatalf("No se esperaba ningún error, se obtuvo %v", err)
	}

	if w.Code != 200 {
		t.Fatalf("Se esperaba código=200, se obtuvo %d", w.Code)
	}

	want := "{\"Personas\":[\"index\",\"out\",\"of\",\"array\"],\"mensaje\":\"Texto analizado con éxito\"}"

	if want != w.Body.String() {
		t.Fatalf("Error obteniendo personas, cadena esperada: %v. Cadena obtenida %v", want, w.Body.String())
	}
}

func TestObtenerEstadisticas(t *testing.T) {
	r := setupServer()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/obtener-estadisticas?posicion=0", nil)

	r.ServeHTTP(w, req)
	//resp, err := http.Get("http://localhost:8080/obtener-redundantes?posicion=0")

	if err != nil {
		t.Fatalf("No se esperaba ningún error, se obtuvo %v", err)
	}

	if w.Code != 200 {
		t.Fatalf("Se esperaba código=200, se obtuvo %d", w.Code)
	}

	want := "{\"Estadísticas\":{\"index out of array\":1},\"mensaje\":\"Texto analizado con éxito\"}"

	if want != w.Body.String() {
		t.Fatalf("Error obteniendo estadisticas, cadena esperada: %v. Cadena obtenida %v", want, w.Body.String())
	}
}
