package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"DemoGraphQL/graph/generated"
	"DemoGraphQL/graph/model"
	"context"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
)

func (r *mutationResolver) CreateGroup(ctx context.Context, input model.NewGroup) (model.Group, error) {
	var group model.Group

	group = model.Group{
		ID:          rand.Int63(),
		Name:        input.Name,
		Description: input.Description,
	}

	r.groups[group.ID] = group
	return group, nil
}

func (r *queryResolver) Groups(ctx context.Context) ([]model.Group, error) {
	var groups []model.Group
	for _, v := range r.groups {
		groups = append(groups, v)
	}
	return groups, nil
}

func (r *queryResolver) GroupTypes(ctx context.Context) ([]model.GroupType, error) {
	jsonFile, err := os.Open("group_types.json")
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	var groupTypes []model.GroupType
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &groupTypes)
	return groupTypes, nil
}

func (r *queryResolver) Items(ctx context.Context) ([]model.Item, error) {
	return r.items, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
