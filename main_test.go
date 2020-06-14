package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func TestIndexHandler(t *testing.T) {
	mux := chi.NewRouter()

	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	mux.Get("/", indexHandler)

	mux.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("expected code 200, got %d", w.Code)
	}

	body := w.Body.String()
	const response = "Hello world"

	if body != response {
		t.Fatalf("expected %s, got %s", response, body)
	}
}
