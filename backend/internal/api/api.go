package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Backend struct {
	DB *gorm.DB
}

func Serve(backend *Backend) error {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

  UsersResource{}.Register(router, backend)

	return router.Run()
}
