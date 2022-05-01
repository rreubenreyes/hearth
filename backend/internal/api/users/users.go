package api

import (
  "github.com/rreubenreyes/hearth/internal/api/backend"
  "github.com/rreubenreyes/hearth/internal/api/resource"
	"github.com/rreubenreyes/hearth/internal/models"
	"github.com/gin-gonic/gin"
)

// UsersResource implements the resource.Modeler interface for the User model.
type UsersResource struct{}

func (UsersResource) Model() interface{} {
	return &models.User{}
}

// Register registers all routes for the users resource.
func (r UsersResource) Register(router *gin.Engine, b *backend.Backend) {
	router.GET("/users/:id", func(c *gin.Context) {
		resource.SimpleGetOne(UsersResource{}, c, b)
	})

	router.POST("/users", func(c *gin.Context) {
		resource.SimpleCreate(UsersResource{}, c, b)
	})
}
