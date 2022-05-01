package api

import (
  "github.com/rreubenreyes/hearth/internal/api/backend"
  "github.com/rreubenreyes/hearth/internal/api/resource"
	"github.com/rreubenreyes/hearth/internal/models"
	"github.com/gin-gonic/gin"
)

// PostsResource implements the resource.Modeler interface for the Post model.
type PostsResource struct{}

func (PostsResource) Model() interface{} {
	return &models.User{}
}

// Register registers all routes for the posts resource.
func (r PostsResource) Register(router *gin.Engine, b *backend.Backend) {
	router.GET("/posts/:id", func(c *gin.Context) {
		resource.SimpleGetOne(PostsResource{}, c, b)
	})

	router.POST("/posts", func(c *gin.Context) {
		resource.SimpleCreate(PostsResource{}, c, b)
	})
}
