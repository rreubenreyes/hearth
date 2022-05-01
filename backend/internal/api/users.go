package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rreubenreyes/hearth/internal/models"
)

type UsersResource struct{}

func (UsersResource) one( c *gin.Context, b *Backend) {
	id := c.Param("id")
	user := &models.User{}

  tx := b.DB.First(user, "id = ?", id)
  if tx.Error != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": tx.Error.Error()})
  }

  // TODO: convert user into map[string]interface{}
	resp, err := json.Marshal(user)
	if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"user": string(resp)})
}

func (UsersResource) create(c *gin.Context, b *Backend) {
	user := &models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

  tx := b.DB.Create(&user)
  if tx.Error != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": tx.Error.Error()})
  }

	resp, err := json.Marshal(user)
	if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"user": string(resp)})
}

func (r UsersResource) Register(router *gin.Engine, b *Backend) {
	router.GET("/users/:id", func(c *gin.Context) {
		r.one(c, b)
	})

	router.POST("/users", func(c *gin.Context) {
		r.create(c, b)
	})
}
