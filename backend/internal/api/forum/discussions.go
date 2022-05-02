package forum

import (
	"github.com/rreubenreyes/hearth/internal/models"
)

// DiscussionsResource implements the resource.Modeler interface for the Discussion model.
type DiscussionsResource struct{}

func (DiscussionsResource) Model() interface{} {
	return &models.Discussion{}
}
