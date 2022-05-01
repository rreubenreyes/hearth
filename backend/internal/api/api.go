package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rreubenreyes/hearth/internal/api/backend"
)

func Serve(b *backend.Backend) error {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	UsersResource{}.Register(router, b)

	return router.Run()
}
