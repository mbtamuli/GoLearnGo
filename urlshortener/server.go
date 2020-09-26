package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.Handle("/foo", new(fooHandler))

	http.HandleFunc("/bar", bar)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

type fooHandler struct {
}

func (h *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, foo"))
}
