// main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	port := flag.String("port", "9000", "server port")
	flag.Parse()

	r := chi.NewRouter()

	r.Get("/", indexHandler)

	log.Fatal(http.ListenAndServe(":"+*port, r))
}
