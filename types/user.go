package types

import (
	"time"

	"github.com/graphql-go/graphql"
)

// UserMaster ユーザー情報
type UserMaster struct {
	ID             uint       `db:"id" json:"id"`
	Name           string     `db:"name" json:"name"`
	Email          string     `db:"email" json:"email"`
	Password       string     `db:"password" json:"password"`
	ProfileImageID int        `db:"profile_image_id" json:"profileImageId"`
	CreatedAt      time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt      time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt      *time.Time `db:"deleted_at" json:"deletedAt"`
}

// UserType ユーザータイプ
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"profileImageId": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
