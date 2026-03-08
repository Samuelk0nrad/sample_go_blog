package blog

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"simple_blog_api/config"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

type Handler struct {
	Env         *config.Env
	MinioClient *minio.Client
}

func RegisterRoutes(rg *gin.RouterGroup, handler Handler) {
	blog := rg.Group("/blog")
	{
		blog.GET("", handler.getPosts)
	}
}

func (h *Handler) getPosts(c *gin.Context) {
	// posts := GetAllPosts()

	reqParams := make(url.Values)
	reqParams.Set("responce-content-disposition", `attachement; filename="TestOBJ"`)

	u, err := h.MinioClient.PresignedGetObject(
		context.Background(),
		h.Env.MinioBucket,
		"TestOBJ",
		15*time.Minute,
		reqParams,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{"data": u.String()})
}
