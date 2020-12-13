package main

import (
	"net/http"
	"testing"
)

func TestObtenerRedundantes(t *testing.T) {
	go setupServer().Run()
	resp, err := http.Get("http://localhost:8080/obtener-redundantes?posicion=0")

	if err != nil {
		t.Fatalf("No se esperaba ningún error, se obtuvo %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Se esperaba código 200, se obtuvo %v", resp.StatusCode)
	}

	_, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Se esperaba que el Content-Type estuviera establecido")
	}
}

func TestObtenerSinRedundantes(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/obtener-sin-redundantes?posicion=0")

	if err != nil {
		t.Fatalf("No se esperaba ningún error, se obtuvo %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Se esperaba código 200, se obtuvo %v", resp.StatusCode)
	}

	_, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Se esperaba que el Content-Type estuviera establecido")
	}
}

func TestObtenerPersonas(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/obtener-personas?posicion=0")

	if err != nil {
		t.Fatalf("No se esperaba ningún error, se obtuvo %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Se esperaba código 200, se obtuvo %v", resp.StatusCode)
	}

	_, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Se esperaba que el Content-Type estuviera establecido")
	}
}

func TestObtenerEstadisticas(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/obtener-estadisticas?posicion=0")

	if err != nil {
		t.Fatalf("No se esperaba ningún error, se obtuvo %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Se esperaba código 200, se obtuvo %v", resp.StatusCode)
	}

	_, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Se esperaba que el Content-Type estuviera establecido")
	}
}
