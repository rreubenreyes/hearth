package api

import (
	"encoding/json"

	"github.com/rreubenreyes/hearth/internal/models"
	"gorm.io/gorm"
)

func Create(db gorm.DB, req []byte) ([]byte, error) {
	var user *models.User
	err := json.Unmarshal(req, user)
	if err != nil {
		return nil, err
	}

	result := db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return json.Marshal(user)
}
