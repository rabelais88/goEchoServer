package controller

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"
)

func Hello(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return "World", nil
		},
		Description: "Hello world",
	}
}

func rootQuery(db *gorm.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"Hello": Hello(db),
		},
	})
}

func GraphQLHandler(db *gorm.DB) (*handler.Handler, error) {

	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery(db)})

	if err != nil {
		return nil, err
	}
	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	}), nil
}
