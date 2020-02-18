package controller

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var PostType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id":      &graphql.Field{Type: graphql.Int},
			"title":   &graphql.Field{Type: graphql.String},
			"author":  &graphql.Field{Type: graphql.String},
			"content": &graphql.Field{Type: graphql.String},
		},
	},
)

func gqlAddPost(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type:        PostType,
		Description: "create a new post",
		Args: graphql.FieldConfigArgument{
			"title":   &graphql.ArgumentConfig{Type: graphql.String},
			"content": &graphql.ArgumentConfig{Type: graphql.String},
			"author":  &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			title, ok := p.Args["title"].(string)
			if ok == false {
				return nil, MakeGraphQLError("NO_TITLE")
			}
			content, ok := p.Args["content"].(string)
			if ok == false {
				return nil, MakeGraphQLError("NO_CONTENT")
			}
			author, ok := p.Args["author"].(string)
			if ok == false {
				return nil, MakeGraphQLError("NO_AUTHOR")
			}
			post := Post{
				Author:  author,
				Content: content,
				Title:   title,
			}
			db.Create(&post)
			return post, nil
		},
	}
}

func gqlGetPost(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type:        PostType,
		Description: "query a post with id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			postId, ok := p.Args["id"].(uint)
			if ok == false {
				return nil, MakeGraphQLError("WRONG_POST_ID")
			}
			post := new(PostModel)
			db.First(&post, postId)
			return post, nil
		},
	}
}
