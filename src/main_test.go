package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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
}
