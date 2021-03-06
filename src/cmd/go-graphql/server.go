package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mkm29/blogql/cmd/go-graphql/graph/model"
	"github.com/mkm29/blogql/pkg/postgres"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mkm29/blogql/cmd/go-graphql/graph"
	"github.com/mkm29/blogql/cmd/go-graphql/graph/generated"
)

const defaultPort = "8080"

var port = os.Getenv("PORT")
var dbUserName = os.Getenv("POSTGRES_USER")
var dbPassword = os.Getenv("POSTGRES_PASSWORD")
var dbURL = os.Getenv("POSTGRES_URL")
var dbName = os.Getenv("POSTGRES_DB")

func main() {
	if port == "" {
		port = defaultPort
	}

	toDoService := &postgres.ToDoImpl{
		DbUserName: dbUserName,
		DbPassword: dbPassword,
		DbURL:      dbURL,
		DbName:     dbName,
	}

	err := toDoService.Initialise()
	if err != nil {
		log.Fatal(err)
	}

	config := generated.Config{Resolvers: &graph.Resolver{ToDo: toDoService}}
	config.Directives.HasRole = hasRoleDirective

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func hasRoleDirective(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (res interface{}, err error) {
	log.Printf("Inside hasRoleDirective - ignore the role check for now")
	return next(ctx)
}
