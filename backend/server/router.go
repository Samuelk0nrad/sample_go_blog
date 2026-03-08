package server

import (
	"simple_blog_api/module/blog"

	"github.com/gin-gonic/gin"
)

func NewRouter(app *App) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		blog.RegisterRoutes(api, blog.Handler{
			Env:         app.Env,
			MinioClient: app.MinioClient,
		})
	}

	return r
}
