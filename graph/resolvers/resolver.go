package resolvers

import (
	gographql "github.com/bcmmbaga/go-graphql"
)

//go:generate go run github.com/99designs/gqlgen --verbose

type Resolver struct {
	Meetups gographql.MeetupRepository
	Users   gographql.UserRepository
}
