package model

import "github.com/go-programing-tour-book/blog-service/pkg/app"

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

type Article struct {
	*Model
	Title         	string 		`json:"title"`
	Desc          	string 		`json:"desc"`
	Content       	string 		`json:"content"`
	CoverImageUrl string 	`json:"cover_image_url"`
	State         	uint8  		`json:"state"`
}
