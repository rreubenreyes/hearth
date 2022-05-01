package users

import (
	"github.com/gin-gonic/gin"
	"github.com/rreubenreyes/hearth/internal/api/backend"
	"github.com/rreubenreyes/hearth/internal/api/resource"
	"github.com/rreubenreyes/hearth/internal/models"
)

// UsersResource implements the resource.Modeler interface for the User model.
type UsersResource struct{}

func (UsersResource) Model() interface{} {
	return &models.User{}
}

// Register registers all routes for the users resource.
func Register(router *gin.Engine, b *backend.Backend) {
	router.GET("/v1/users/:id", func(c *gin.Context) {
		resource.SimpleGetOne(UsersResource{}, c, b)
	})

	router.POST("/v1/users", func(c *gin.Context) {
		resource.SimpleCreate(UsersResource{}, c, b)
	})
}
