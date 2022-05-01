package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rreubenreyes/hearth/internal/api/backend"
	"github.com/rreubenreyes/hearth/internal/api/forum"
	"github.com/rreubenreyes/hearth/internal/api/users"
)

func Serve(b *backend.Backend) error {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	users.Register(router, b)
	forum.Register(router, b)

	return router.Run()
}
