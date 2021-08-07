package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"practice_gql/gqlconv"
	"practice_gql/graphql/graph/generated"
	"practice_gql/graphql/graph/model"
	"practice_gql/interfaces/database"
)

func (r *mutationResolver) CreateBlog(ctx context.Context, input model.NewBlog) (*model.Blog, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Blogs(ctx context.Context) ([]*model.Blog, error) {
	// panic(fmt.Errorf("not implemented"))
	blogs, err := database.BlogAll(r.DB)
	if err != nil {
		log.Println(err)
	}
	out := gqlconv.Blogs(blogs)
	return out, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
