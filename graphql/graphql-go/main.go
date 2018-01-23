package main

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var data map[string]user = map[string]user{
	"1": user{
		ID:   "1",
		Name: "Dan",
	},
	"2": user{
		ID:   "2",
		Name: "Lee",
	},
	"3": user{
		ID:   "3",
		Name: "Nick",
	},
}

//   Create User object type with fields "id" and "name" by using GraphQLObjectTypeConfig:
//       - Name: name of object type
//       - Fields: a map of fields by using GraphQLFields
//   Setup type of field use GraphQLFieldConfig
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

//   Create Query object type with fields "user" has type [userType] by using GraphQLObjectTypeConfig:
//       - Name: name of object type
//       - Fields: a map of fields by using GraphQLFields
//   Setup type of field use GraphQLFieldConfig to define:
//       - Type: type of field
//       - Args: arguments to query with current field
//       - Resolve: function to query data using params from [Args] and return value with current type
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						if v, exist := data[idQuery]; exist {
							return v, nil
						}

						return nil, nil
					}
					return nil, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema, vars map[string]interface{}) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  query,
		VariableValues: vars,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	query := `query userinfo($uid: String = "1") {
				user(id: $uid){
					id
					name
				}
			}
		`
	vars := map[string]interface{}{"uid": "3"}
	r := executeQuery(query, schema, vars)
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
}
