package test

import "time"

//go:generate go run ../../main.go -in models.go -out autogenerated_queryset.go

type model struct {
	ID                   int
	DeletedAt, UpdatedAt time.Time
}

// User is a usual user
// gen:qs
type User struct {
	model

	Posts []Post
	Name  string
}

// Blog is a blog
// gen:qs
type Blog struct {
	model

	Name string
}

// Post is an article
// gen:qs
type Post struct {
	model

	Blog  *Blog // may be no blog
	User  User
	Title string
}