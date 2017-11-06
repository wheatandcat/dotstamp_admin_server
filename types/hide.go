package types

import (
	"github.com/graphql-go/graphql"
)

// Hide 非表示
type Hide struct {
	ID uint `json:"id"`
}

// HideType 　非表示タイプ
var HideType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Hide",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "id",
			},
		},
	},
)
