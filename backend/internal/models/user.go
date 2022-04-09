package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `gorm:"->;<-:create;primaryKey;unique"`
	Name        string    `gorm:"->;<-:create"`
	DisplayName string
	Profile     Profile   `gorm:"embedded"`
	CreatedAt   time.Time `gorm:"->;<-:create;autoCreateTime:milli"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:milli"`
}

type Profile struct {
	Age                     int
	Bio                     string
	Birthday                time.Time
	DiscussionsInterestedIn []*Discussion `gorm:"many2many:user_interested_discussions"`
	Friends                 []*User       `gorm:"many2many:user_friends"`
}
