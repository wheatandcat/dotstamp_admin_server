package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	yaml "gopkg.in/yaml.v2"

	_ "github.com/go-sql-driver/mysql"
	"github.com/graphql-go/graphql"
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
						Type: graphql.Int,
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
				Description: "userList",
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type: graphql.Int,
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
						Type: graphql.Int,
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
				Description: "find contribution list",
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					first, _ := p.Args["first"].(int)
					u := []types.UserContribution{}
					err := DB.Select(&u, "SELECT id,title,view_status FROM user_contributions ORDER BY id ASC LIMIT ?", first)
					if err != nil {
						return nil, nil
					}

					return u, nil
				},
			},
			"problemList": &graphql.Field{
				Type:        graphql.NewList(types.LogProblemContributionReportType),
				Description: "find proble list",
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type: graphql.Int,
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
				Description: "find question list",
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type: graphql.Int,
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

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	connectDB()

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={user(id:1){name}}'")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={userList(first:100){id,name}}'")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={contribution(id:1){id,title}}'")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={contributionList(first:100){id,title}}'")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={problemList(first:100){id}}'")

	http.ListenAndServe(":8080", nil)
}

//Helper function to import json from file to map
func importJSONDataFromFile(fileName string, result interface{}) (isOK bool) {
	isOK = true
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print("Error:", err)
		isOK = false
	}
	err = json.Unmarshal(content, result)
	if err != nil {
		isOK = false
		fmt.Print("Error:", err)
	}
	return
}
