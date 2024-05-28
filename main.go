package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gustavopcr/touchdown/api"
	"github.com/gustavopcr/touchdown/graph"
	"github.com/gustavopcr/touchdown/internal"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	internal.StartTouchdown()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/query", srv)
	http.HandleFunc("/verify", api.HandleVerify)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
