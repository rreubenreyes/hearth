package forum

import (
	"github.com/rreubenreyes/hearth/internal/models"
)

// BoardsResource implements the resource.Modeler interface for the Board model.
type BoardsResource struct{}

func (BoardsResource) Model() interface{} {
	return &models.Board{}
}
