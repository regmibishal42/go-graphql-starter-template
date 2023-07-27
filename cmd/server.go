package main

import (
	"backend/graph/generated"
	"backend/graph/resolvers"
	"backend/infrastructure/db"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	dbClient := db.InitDB()
	fmt.Println(dbClient)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	resolver := resolvers.NewResolver()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	// srv.AddTransport(&transport.Websocket{})
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
