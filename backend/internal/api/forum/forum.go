package forum

import (
	"github.com/gin-gonic/gin"
	"github.com/rreubenreyes/hearth/internal/api/backend"
	"github.com/rreubenreyes/hearth/internal/api/resource"
)

// Register registers all routes for the forum resource.
func Register(router *gin.Engine, b *backend.Backend) {
  // posts
	router.GET("/v1/forum/posts/:id", func(c *gin.Context) {
		resource.SimpleGetOne(PostsResource{}, c, b)
	})

	router.POST("/v1/forum/posts", func(c *gin.Context) {
		resource.SimpleCreate(PostsResource{}, c, b)
	})

  // discussions
	router.GET("/v1/forum/discussions/:id", func(c *gin.Context) {
		resource.SimpleGetOne(DiscussionsResource{}, c, b)
	})

	router.POST("/v1/forum/discussions", func(c *gin.Context) {
		resource.SimpleCreate(DiscussionsResource{}, c, b)
	})

  // boards
	router.GET("/v1/forum/boards/:id", func(c *gin.Context) {
		resource.SimpleGetOne(BoardsResource{}, c, b)
	})

	router.POST("/v1/forum/boards", func(c *gin.Context) {
		resource.SimpleCreate(BoardsResource{}, c, b)
	})
}
