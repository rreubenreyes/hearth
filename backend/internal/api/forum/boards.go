package api

import (
  "github.com/rreubenreyes/hearth/internal/api/backend"
  "github.com/rreubenreyes/hearth/internal/api/resource"
	"github.com/rreubenreyes/hearth/internal/models"
	"github.com/gin-gonic/gin"
)

// BoardsResource implements the resource.Modeler interface for the Board model.
type BoardsResource struct{}

func (BoardsResource) Model() interface{} {
	return &models.User{}
}

// Register registers all routes for the boards resource.
func (r BoardsResource) Register(router *gin.Engine, b *backend.Backend) {
	router.GET("/boards/:id", func(c *gin.Context) {
		resource.SimpleGetOne(BoardsResource{}, c, b)
	})

	router.POST("/boards", func(c *gin.Context) {
		resource.SimpleCreate(BoardsResource{}, c, b)
	})
}
