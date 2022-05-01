// The resource package implements a generic CRUD interface for a given model.
// The models and types contained within are meant to be sufficient for simple
// CRUD operations without intermediate functionality.
package resource

import (
	"net/http"

  "github.com/rreubenreyes/hearth/internal/api/backend"
	"github.com/gin-gonic/gin"
)

// Modeler is an interface which exposes a Model() method. The Model() method
// returns a pointer to a model which can be used to interact with the database.
type Modeler interface {
	Model() interface{}
}

// SimpleGetOne is a generic implementation of a GET request for a single resource.
// The first argument implements the Modeler interface.
func SimpleGetOne(m Modeler, c *gin.Context, b *backend.Backend) {
	id := c.Param("id")
	entity := m.Model()

	tx := b.DB.First(entity, "id = ?", id)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": tx.Error.Error()})
	}

	c.JSON(http.StatusOK, entity)
}

// SimpleGetOne is a generic implementation of a POST request for a single resource.
// The first argument implements the Modeler interface.
func SimpleCreate(m Modeler, c *gin.Context, b *backend.Backend) {
	entity := m.Model()
	err := c.ShouldBindJSON(entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	tx := b.DB.Create(entity)
	if tx.Error != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": tx.Error.Error()})
	}

	c.JSON(http.StatusCreated, entity)
}
