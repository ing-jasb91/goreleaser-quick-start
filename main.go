// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	r := chi.NewRouter()

	r.Get("/", indexHandler)

	log.Fatal(http.ListenAndServe(":9000", r))
}
