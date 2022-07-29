package main

import graphql "github.com/graph-gophers/graphql-go"

var (
	opts  = []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	posts = []Post{
		{ID: "0", Title: "Handy Golang tutorials", Body: "Here's the link to the videos: https://www.youtube.com/watch?v=etSN4X_fCnM&list=PL4cUxeGkcC9gC88BEo9czgyS72A3doDeM", PostedBy: users[0]},
		{ID: "1", Title: "GraphQL documentation", Body: "The official documentation for GraphQL can be found at https://graphql.org/", PostedBy: users[1]},
	}
	users = []User{
		{ID: "0", Name: "Adam", Email: "adam@email.com"},
		{ID: "1", Name: "Jack", Email: "jack@hotmail.co.uk"},
		{ID: "2", Name: "Sol", Email: "sol@atos.net"},
	}
)

type User struct {
	ID    graphql.ID
	Name  string
	Email string
	Posts []Post
}

type Post struct {
	ID       graphql.ID
	Title    string
	Body     string
	PostedBy User
}

type RootResolver struct{}
