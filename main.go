package main

import (
	"net/http"

	"github.com/graphql-go/handler"
)

func main() {
	h := handler.New(&handler.Config{
		Schema:   &Schema,
		GraphiQL: true,
		Pretty:   true,
	})

	http.Handle("/graphql", h)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It works!"))
	})
	http.ListenAndServe(":8000", nil)
}
