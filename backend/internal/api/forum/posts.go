package forum

import (
	"github.com/rreubenreyes/hearth/internal/models"
)

// PostsResource implements the resource.Modeler interface for the Post model.
type PostsResource struct{}

func (PostsResource) Model() interface{} {
	return &models.User{}
}
