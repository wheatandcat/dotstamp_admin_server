package types

import (
	"github.com/graphql-go/graphql"
)

// UserContribution ユーザ投稿
type UserContribution struct {
	ID         uint    `db:"id" json:"id"`
	UserID     int     `db:"user_id" json:"userId"`
	Title      string  `db:"title" json:"title"`
	ViewStatus int     `db:"view_status" json:"viewStatus"`
	CreatedAt  string  `db:"created_at" json:"createdAt"`
	UpdatedAt  string  `db:"updated_at" json:"updatedAt"`
	DeletedAt  *string `db:"deleted_at" json:"deletedAt"`
}

// ContributionType 　投稿タイプ
var ContributionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Contribution",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "id",
			},
			"userId": &graphql.Field{
				Type:        graphql.Int,
				Description: "user id",
			},
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "title",
			},
			"viewStatus": &graphql.Field{
				Type:        graphql.Int,
				Description: "view status 0 ore 1",
			},
			"createdAt": &graphql.Field{
				Type:        graphql.String,
				Description: "created date",
			},
			"updatedAt": &graphql.Field{
				Type:        graphql.String,
				Description: "update date",
			},
			"deletedAt": &graphql.Field{
				Type:        graphql.String,
				Description: "delete date",
			},
		},
	},
)
