package backend

import (
	"gorm.io/gorm"
)

type Backend struct {
	DB *gorm.DB
}
