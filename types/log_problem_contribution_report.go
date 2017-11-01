package types

import (
	"github.com/graphql-go/graphql"
)

// LogProblemContributionReport 投稿通報ログ
type LogProblemContributionReport struct {
	ID                 uint    `db:"id" json:"id"`
	UserID             int     `db:"user_id" json:"userId"`
	Type               int     `db:"type" json:"type"`
	UserContributionID int     `db:"user_contribution_id" json:"userContributionId"`
	CreatedAt          string  `db:"created_at" json:"createdAt"`
	UpdatedAt          string  `db:"updated_at" json:"updatedAt"`
	DeletedAt          *string `db:"deleted_at" json:"deletedAt"`
}

// LogProblemContributionReportType 投稿通報ログタイプ
var LogProblemContributionReportType = graphql.NewObject(
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
			"type": &graphql.Field{
				Type:        graphql.Int,
				Description: "title",
			},
			"userContributionId": &graphql.Field{
				Type:        graphql.Int,
				Description: "contribution id",
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
