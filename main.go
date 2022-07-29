package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go/relay"
)

func main() {

	http.Handle("/graphql", CorsMiddleware(&relay.Handler{
		Schema: parseSchema("./schema.graphql", &RootResolver{}),
	}))

	fmt.Println("serving on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
