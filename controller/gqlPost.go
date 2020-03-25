package controller

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var PostType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id":        &graphql.Field{Type: graphql.Int},
			"title":     &graphql.Field{Type: graphql.String},
			"author":    &graphql.Field{Type: graphql.String},
			"content":   &graphql.Field{Type: graphql.String},
			"createdAt": &graphql.Field{Type: graphql.DateTime},
			"updatedAt": &graphql.Field{Type: graphql.DateTime},
			"deletedAt": &graphql.Field{Type: graphql.DateTime},
		},
	},
)

var PostsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Posts",
	Fields: graphql.Fields{
		"items": &graphql.Field{Type: PostType},
		"page":  &graphql.Field{Type: graphql.Int},
		"count": &graphql.Field{Type: graphql.Int},
	},
})

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

func gqlGetPosts(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type:        PostsType,
		Description: "query multiple posts",
		Args: graphql.FieldConfigArgument{
			"page": &graphql.ArgumentConfig{Type: graphql.Int},
			"size": &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			page, ok := p.Args["page"].(int)
			if ok == false {
				return nil, MakeGraphQLError("WRONG_PAGE")
			}
			size, ok := p.Args["size"].(int)
			if ok == false {
				return nil, MakeGraphQLError("WRONG_SIZE")
			}
			var count uint
			posts := []PostModel{}
			currentPage := page * size
			db.Offset(uint(currentPage)).Find(&posts).Count(&count)
			result := PostsResponse{
				Items: posts,
				Count: count,
				Page:  uint(currentPage),
			}
			return result, nil
		},
	}
}
