package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bcmmbaga/go-graphql/graph/generated"
	"github.com/bcmmbaga/go-graphql/models"
)

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Meeutps(ctx context.Context) ([]*models.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type meetupResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
