package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `gorm:"->;<-:create;primaryKey;unique" json:"id"`
	Name        string    `gorm:"->;<-:create" json:"name"`
	DisplayName string    `json:"display_name"`
	Profile     Profile   `gorm:"embedded" json:"profile"`
	CreatedAt   time.Time `gorm:"->;<-:create;autoCreateTime:milli"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:milli"`
}

type Profile struct {
	Age                     int           `json:"age"`
	Bio                     string        `json:"bio"`
	Birthday                time.Time     `json:"birthday"`
	DiscussionsInterestedIn []*Discussion `gorm:"many2many:user_interested_discussions" json:"discussions_interested_in"`
	Friends                 []*User       `gorm:"many2many:user_friends" json:"friends"`
}
