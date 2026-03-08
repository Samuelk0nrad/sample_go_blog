package service

import "simple_blog_api/internal/model"

func GetAllPosts() []model.BlogPost {
	return []model.BlogPost{
		{ID: 1, Title: "First blog post"},
		{ID: 2, Title: "Second blog post"},
	}
}
