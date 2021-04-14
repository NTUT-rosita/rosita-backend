package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/NTUT-rosita/rosita-backend/internal/api/graphql"
	"github.com/NTUT-rosita/rosita-backend/internal/api/graphql/generated"
	"github.com/NTUT-rosita/rosita-backend/internal/db"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbClient, err := sql.Open("postgres", "")
	if err != nil {
		panic("db connect error:" + err.Error())
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{Queries: *db.New(dbClient)}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
