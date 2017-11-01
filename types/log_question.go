package types

import (
	"github.com/graphql-go/graphql"
)

// LogQuestion 質問ログ
type LogQuestion struct {
	ID        uint    `db:"id" json:"id"`
	UserID    int     `db:"user_id" json:"userId"`
	Email     string  `db:"email" json:"email"`
	Body      string  `db:"body" json:"body"`
	CreatedAt string  `db:"created_at" json:"createdAt"`
	UpdatedAt string  `db:"updated_at" json:"updatedAt"`
	DeletedAt *string `db:"deleted_at" json:"deletedAt"`
}

// LogQuestionType 質問ログタイプ
var LogQuestionType = graphql.NewObject(
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
			"email": &graphql.Field{
				Type:        graphql.String,
				Description: "email",
			},
			"body": &graphql.Field{
				Type:        graphql.String,
				Description: "body",
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
