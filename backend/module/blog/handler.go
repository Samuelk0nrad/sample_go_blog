package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	blog := rg.Group("/blog")
	{
		blog.GET("", getPosts)
	}
}

func getPosts(c *gin.Context) {
	posts := GetAllPosts()
	c.JSON(http.StatusOK, gin.H{"data": posts})
}
