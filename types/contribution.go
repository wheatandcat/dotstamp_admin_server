package types

import (
	"time"

	"github.com/graphql-go/graphql"
)

// UserContribution ユーザ投稿
type UserContribution struct {
	ID         uint       `db:"id" json:"id"`
	UserID     int        `db:"user_id" json:"userId"`
	Title      string     `db:"title" json:"title"`
	ViewStatus int        `db:"view_status" json:"viewStatus"`
	CreatedAt  time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt  time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt  *time.Time `db:"deleted_at" json:"deletedAt"`
}

// ContributionType 　投稿タイプ
var ContributionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Contribution",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"viewStatus": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
