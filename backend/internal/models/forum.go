package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Post is an archetypal forum post.
type Post struct {
	gorm.Model
	ID        uuid.UUID `gorm:"->;<-:create;primaryKey;unique"`
	User      *User     `gorm:"->;<-:create"`
	CreatedAt time.Time `gorm:"->;<-:create;autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}

// Discussion is a typical representation of a forum thread.
type Discussion struct {
	gorm.Model
	ID              uuid.UUID `gorm:"->;<-:create;primaryKey;unique"`
	Title           string
	Posts           []*Post
	Views           int
	InterestedUsers []*User   `gorm:"many2many:user_interested_discussions"`
	CreatedAt       time.Time `gorm:"->;<-:create;autoCreateTime:milli"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime:milli"`
}

// Board is a collection of discussions.
type Board struct {
	gorm.Model
	ID          uuid.UUID `gorm:"->;<-:create;primaryKey;unique"`
	Name        string
	Discussions []*Discussion
	CreatedAt   time.Time `gorm:"->;<-:create;autoCreateTime:milli"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:milli"`
}
