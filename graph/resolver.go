//go:generate go run github.com/99designs/gqlgen generate
package graph

import "DemoGraphQL/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	groups map[int64]model.Group
	items  []model.Item
}
