package main

import (
	"log"
	"net/http"
	"os"
	"practice_gql/graphql/graph"
	"practice_gql/graphql/graph/generated"
	"practice_gql/interfaces/database"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	/**  additional  **/
	db, err := database.ConnectSql()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	/**  end of additional  **/

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
