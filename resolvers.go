package main

import (
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
)

func (r *RootResolver) AllPosts() ([]Post, error) {
	return posts, nil
}

func (r *RootResolver) AllUsers() ([]User, error) {
	return users, nil
}

func (r *RootResolver) CreatePost(args struct {
	Title string
	Body  string
}) (Post, error) {
	newPost := Post{
		ID:       graphql.ID(fmt.Sprint(len(posts))),
		Title:    args.Title,
		Body:     args.Body,
		PostedBy: users[2],
	}
	posts = append(posts, newPost)
	return newPost, nil
}

func (r *RootResolver) DeletePost(args struct {
	ID int32
}) (Post, error) {
	post := posts[args.ID]
	// posts = RemoveIndex(posts[args.ID])
	posts = append(posts[:args.ID], posts[args.ID+1:]...)
	return post, nil
}
