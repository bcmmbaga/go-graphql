package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bcmmbaga/go-graphql/graph/generated"
	"github.com/bcmmbaga/go-graphql/graph/resolvers"
	"github.com/bcmmbaga/go-graphql/postgres"
	"github.com/sirupsen/logrus"
)

const defaultPort = "8080"

func main() {
	db := postgres.Connect("172.17.0.2:5432", "test", "test", "test")
	if db == nil {
		logrus.Fatalln("failed to setup database")
	}

	defer db.Close()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("\nconnect to http://localhost:%s/ for GraphQL playground\n", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
