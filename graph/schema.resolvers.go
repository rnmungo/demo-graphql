package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"DemoGraphQL/graph/generated"
	"DemoGraphQL/graph/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

func (r *mutationResolver) CreateGroup(ctx context.Context, input model.NewGroup) (*model.Group, error) {
	jsonBytes, err := readJson("group_types.json")
	if err != nil {
		return nil, err
	}

	var groupTypes []*model.GroupType
	json.Unmarshal(jsonBytes, &groupTypes)

	var groupType *model.GroupType
	for _, v := range groupTypes {
		if v.ID == input.GroupTypeID {
			groupType = v
		}
	}

	var group *model.Group
	group = &model.Group{
		ID:          fmt.Sprintf("%d", rand.Int63()),
		Name:        input.Name,
		Description: input.Description,
		GroupType:   groupType,
	}

	r.GroupsStore[group.ID] = group
	return group, nil
}

func (r *mutationResolver) CreateItem(ctx context.Context, input []*model.NewItem) (*model.Item, error) {
	body := input[0]
	group, ok := r.GroupsStore[body.GroupID]
	if !ok {
		return &model.Item{}, errors.New("group not found")
	}

	item := &model.Item{
		Value: body.Value,
		Group: group,
	}

	r.itemsStore = append(r.itemsStore, item)
	return item, nil
}

func (r *mutationResolver) RemoveItem(ctx context.Context, input []*model.NewItem) (*int, error) {
	var rowsAffected int
	body := input[0]
	for i, v := range r.itemsStore {
		if v.Value == body.Value && v.Group.ID == body.GroupID {
			r.itemsStore = remove(r.itemsStore, i)
			rowsAffected++
		}
	}

	return &rowsAffected, nil
}

func (r *queryResolver) GroupTypes(ctx context.Context) ([]*model.GroupType, error) {
	jsonBytes, err := readJson("group_types.json")
	if err != nil {
		return nil, err
	}

	var groupTypes []*model.GroupType
	json.Unmarshal(jsonBytes, &groupTypes)
	return groupTypes, nil
}

func (r *queryResolver) Groups(ctx context.Context) ([]*model.Group, error) {
	var groups []*model.Group
	for _, v := range r.GroupsStore {
		groups = append(groups, v)
	}
	return groups, nil
}

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	return r.itemsStore, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func readJson(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	return ioutil.ReadAll(jsonFile)
}

func remove(s []*model.Item, i int) []*model.Item {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
