package api

import (
  "github.com/rreubenreyes/hearth/internal/api/backend"
  "github.com/rreubenreyes/hearth/internal/api/resource"
	"github.com/rreubenreyes/hearth/internal/models"
	"github.com/gin-gonic/gin"
)

// DiscussionsResource implements the resource.Modeler interface for the Discussion model.
type DiscussionsResource struct{}

func (DiscussionsResource) Model() interface{} {
	return &models.User{}
}

// Register registers all routes for the discussions resource.
func (r DiscussionsResource) Register(router *gin.Engine, b *backend.Backend) {
	router.GET("/discussions/:id", func(c *gin.Context) {
		resource.SimpleGetOne(DiscussionsResource{}, c, b)
	})

	router.POST("/discussions", func(c *gin.Context) {
		resource.SimpleCreate(DiscussionsResource{}, c, b)
	})
}
