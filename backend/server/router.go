package server

import (
	"simple_blog_api/module/blog"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		blog.RegisterRoutes(api)
	}

	return r
}
