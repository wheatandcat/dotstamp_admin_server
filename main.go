package main

import (
	"io/ioutil"
	"log"
	"net/http"

	yaml "gopkg.in/yaml.v2"

	_ "github.com/go-sql-driver/mysql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jmoiron/sqlx"
	"github.com/wheatandcat/dotstamp_admin_server/types"
)

// DB database connection
var DB *sqlx.DB

// DbInfo DB情報
type DbInfo struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

func connectDB() {
	buf, err := ioutil.ReadFile("./config/devlop.yaml")
	if err != nil {
		panic(err)
	}

	var d DbInfo
	err = yaml.Unmarshal(buf, &d)
	if err != nil {
		panic(err)
	}

	db, err := sqlx.Connect("mysql", d.User+d.Password+":@/"+d.Dbname)
	log.Println(d.User + d.Password + ":@/" + d.Dbname)
	if err != nil {
		panic(err)
	}

	DB = db
}

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type:        types.UserType,
				Description: "find user",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:        graphql.Int,
						Description: "user id",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(int)
					u := types.UserMaster{}
					err := DB.Get(&u, "SELECT id,name,email,password,profile_image_id FROM user_masters WHERE id=?", id)
					if err != nil {
						return nil, nil
					}

					return u, nil
				},
			},
			"userList": &graphql.Field{
				Type:        graphql.NewList(types.UserType),
				Description: "user list",
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type:        graphql.Int,
						Description: "number of item displayed",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					first, _ := p.Args["first"].(int)
					u := []types.UserMaster{}
					err := DB.Select(&u, "SELECT id,name,email,password,profile_image_id FROM user_masters ORDER BY id ASC LIMIT ?", first)
					if err != nil {
						return nil, nil
					}

					return u, nil
				},
			},
			"contribution": &graphql.Field{
				Type:        types.ContributionType,
				Description: "find contribution",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:        graphql.Int,
						Description: "contribution id",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(int)
					u := types.UserContribution{}
					err := DB.Get(&u, "SELECT id,title,view_status FROM user_contributions WHERE id=?", idQuery)
					if err != nil {
						return nil, nil
					}

					return u, nil
				},
			},
			"contributionList": &graphql.Field{
				Type:        graphql.NewList(types.ContributionType),
				Description: "contribution list",
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type:        graphql.Int,
						Description: "number of item displayed",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					first, _ := p.Args["first"].(int)
					u := []types.UserContribution{}
					err := DB.Select(&u, "SELECT id,user_id,title,view_status,created_at,updated_at,deleted_at FROM user_contributions ORDER BY id ASC LIMIT ?", first)
					if err != nil {
						return nil, nil
					}

					return u, nil
				},
			},
			"problemList": &graphql.Field{
				Type:        graphql.NewList(types.LogProblemContributionReportType),
				Description: "proble list",
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type:        graphql.Int,
						Description: "number of item displayed",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					first, _ := p.Args["first"].(int)
					r := []types.LogProblemContributionReport{}
					err := DB.Select(&r, "SELECT id,user_contribution_id,user_id,type FROM log_problem_contribution_reports ORDER BY id ASC LIMIT ?", first)
					if err != nil {
						return nil, nil
					}

					return r, nil
				},
			},

			"questionList": &graphql.Field{
				Type:        graphql.NewList(types.LogQuestionType),
				Description: "question list",
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type:        graphql.Int,
						Description: "number of item displayed",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					first, _ := p.Args["first"].(int)
					r := []types.LogQuestion{}
					err := DB.Select(&r, "SELECT id,user_id,email,body FROM log_questions ORDER BY id ASC LIMIT ?", first)
					if err != nil {
						return nil, nil
					}

					return r, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func main() {
	connectDB()

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
