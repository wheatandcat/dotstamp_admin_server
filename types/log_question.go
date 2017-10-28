package types

import (
	"time"

	"github.com/graphql-go/graphql"
)

// LogQuestion 質問ログ
type LogQuestion struct {
	ID        uint       `db:"id" json:"id"`
	UserID    int        `db:"user_id" json:"userId"`
	Email     string     `db:"email" json:"email"`
	Body      string     `db:"body" json:"body"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt"`
}

// LogQuestionType 質問ログタイプ
var LogQuestionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Contribution",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"userId": &graphql.Field{
				Type: graphql.Int,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"body": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
