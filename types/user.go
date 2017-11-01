package types

import (
	"github.com/graphql-go/graphql"
)

// UserMaster ユーザー情報
type UserMaster struct {
	ID             uint    `db:"id" json:"id"`
	Name           string  `db:"name" json:"name"`
	Email          string  `db:"email" json:"email"`
	Password       string  `db:"password" json:"password"`
	ProfileImageID int     `db:"profile_image_id" json:"profileImageId"`
	CreatedAt      string  `db:"created_at" json:"createdAt"`
	UpdatedAt      string  `db:"updated_at" json:"updatedAt"`
	DeletedAt      *string `db:"deleted_at" json:"deletedAt"`
}

// UserType ユーザータイプ
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "id",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "name",
			},
			"email": &graphql.Field{
				Type:        graphql.String,
				Description: "email",
			},
			"profileImageId": &graphql.Field{
				Type:        graphql.Int,
				Description: "profile image id",
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
