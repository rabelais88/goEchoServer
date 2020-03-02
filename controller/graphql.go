package controller

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"
)

type GraphQLError struct {
	Message string
}

func (e *GraphQLError) Error() string {
	return e.Message
}

func MakeGraphQLError(message string) error {
	return &GraphQLError{Message: message}
}

func Hello(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return "World", nil
		},
		Description: "Hello world!!",
	}
}

func rootQuery(db *gorm.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"Hello": Hello(db),
			"Post":  gqlGetPost(db),
		},
	})
}

func rootMutation(db *gorm.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createPost": gqlAddPost(db),
		},
	})
}

func GetGraphQLSchema(db *gorm.DB) (*graphql.Schema, error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery(db), Mutation: rootMutation(db)})
	if err != nil {
		return nil, err
	}
	return &schema, nil
}

func GraphQLHandler(db *gorm.DB) (*handler.Handler, error) {
	schema, err := GetGraphQLSchema(db)

	if err != nil {
		return nil, err
	}
	return handler.New(&handler.Config{
		Schema:   schema,
		Pretty:   true,
		GraphiQL: true,
	}), nil
}
