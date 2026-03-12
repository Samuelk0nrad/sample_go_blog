package service

import (
	"context"

	"simple_blog_api/internal/db"
	"simple_blog_api/internal/dto"
)

type BlogService struct {
	queries *db.Queries
}

func (s *BlogService) GetBlogs() ([]dto.GetBlogDTO, error) {
	DBBlogs, err := s.queries.ListBlogs(context.Background())
	if err != nil {
		return nil, err
	}
	var blogs []dto.GetBlogDTO

	for _, DBBlog := range DBBlogs {
		blogs = append(blogs, dto.GetBlogDTO{
			ID:     int(DBBlog.ID),
			Author: DBBlog.Author,
			Title:  DBBlog.Title,
			Slug:   DBBlog.Slug,
		})
	}
	return blogs, nil
}
