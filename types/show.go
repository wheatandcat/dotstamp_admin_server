package types

import (
	"github.com/graphql-go/graphql"
)

// Show 表示
type Show struct {
	ID uint `json:"id"`
}

// ShowType 　表示タイプ
var ShowType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Show",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "id",
			},
		},
	},
)
