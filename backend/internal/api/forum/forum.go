package forum

import (
	"github.com/gin-gonic/gin"
	"github.com/rreubenreyes/hearth/internal/api/backend"
	"github.com/rreubenreyes/hearth/internal/api/resource"
)

// Register registers all routes for the forum resource.
func Register(router *gin.Engine, b *backend.Backend) {
	router.GET("/v1/posts/:id", func(c *gin.Context) {
		resource.SimpleGetOne(PostsResource{}, c, b)
	})

	router.POST("/v1/posts", func(c *gin.Context) {
		resource.SimpleCreate(PostsResource{}, c, b)
	})

	router.GET("/v1/discussions/:id", func(c *gin.Context) {
		resource.SimpleGetOne(DiscussionsResource{}, c, b)
	})

	router.POST("/v1/discussions", func(c *gin.Context) {
		resource.SimpleCreate(DiscussionsResource{}, c, b)
	})

	router.GET("/v1/boards/:id", func(c *gin.Context) {
		resource.SimpleGetOne(BoardsResource{}, c, b)
	})

	router.POST("/v1/boards", func(c *gin.Context) {
		resource.SimpleCreate(BoardsResource{}, c, b)
	})
}
