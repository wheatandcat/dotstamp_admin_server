package types

import (
	"time"

	"github.com/graphql-go/graphql"
)

// LogProblemContributionReport 投稿通報ログ
type LogProblemContributionReport struct {
	ID                 uint       `db:"id" json:"id"`
	UserID             int        `db:"user_id" json:"userId"`
	Type               int        `db:"type" json:"type"`
	UserContributionID int        `db:"user_contribution_id" json:"userContributionId"`
	CreatedAt          time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt          time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt          *time.Time `db:"deleted_at" json:"deletedAt"`
}

// LogProblemContributionReportType 投稿通報ログタイプ
var LogProblemContributionReportType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Contribution",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"userId": &graphql.Field{
				Type: graphql.Int,
			},
			"type": &graphql.Field{
				Type: graphql.Int,
			},
			"userContributionId": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
