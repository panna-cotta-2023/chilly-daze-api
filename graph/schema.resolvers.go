package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"chilly_daze_gateway/graph/model"
	"context"
	"fmt"
)

// RegisterUser is the resolver for the registerUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, input *model.RegisterUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: RegisterUser - registerUser"))
}

// StartChill is the resolver for the startChill field.
func (r *mutationResolver) StartChill(ctx context.Context, input model.StartChillInput) (*model.Chill, error) {
	chill, err := r.Srv.AddChill(ctx, input)
	if err != nil {
		return nil, err
	}

	return chill, nil
}

// AddTracePoints is the resolver for the addTracePoints field.
func (r *mutationResolver) AddTracePoints(ctx context.Context, input model.TracePointsInput) ([]*model.TracePoint, error) {
	tracePoints, err := r.Srv.AddTracePoints(ctx, input)
	if err != nil {
		return nil, err
	}

	return tracePoints, nil
}

// AddPhotos is the resolver for the addPhotos field.
func (r *mutationResolver) AddPhotos(ctx context.Context, input model.PhotosInput) ([]*model.Photo, error) {
	photos, err := r.Srv.AddPhotos(ctx, input)
	if err != nil {
		return nil, err
	}

	return photos, nil
}

// EndChill is the resolver for the endChill field.
func (r *mutationResolver) EndChill(ctx context.Context, input model.EndChillInput) (*model.Chill, error) {
	panic(fmt.Errorf("not implemented: EndChill - endChill"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Achievements is the resolver for the achievements field.
func (r *queryResolver) Achievements(ctx context.Context) ([]*model.Achievement, error) {
	panic(fmt.Errorf("not implemented: Achievements - achievements"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
