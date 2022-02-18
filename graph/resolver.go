//go:generate go run github.com/99designs/gqlgen generate
package graph

import "DemoGraphQL/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	GroupsStore map[string]*model.Group
	itemsStore  []*model.Item
}
