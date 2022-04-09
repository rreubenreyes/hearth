package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Post is an archetypal forum post.
type Post struct {
	gorm.Model
	ID        uuid.UUID `gorm:"->;<-:create;primaryKey;unique" json:"id"`
	User      *User     `gorm:"->;<-:create" json:"user"`
	Content   string    `json:"content"`
	CreatedAt time.Time `gorm:"->;<-:create;autoCreateTime:milli" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

// Discussion is a typical representation of a forum thread.
type Discussion struct {
	gorm.Model
	ID              uuid.UUID `gorm:"->;<-:create;primaryKey;unique" json:"id"`
	Title           string    `json:"title"`
	Posts           []*Post   `json:"posts"`
	Views           int       `json:"views"`
	InterestedUsers []*User   `gorm:"many2many:user_interested_discussions" json:"interested_users"`
	CreatedAt       time.Time `gorm:"->;<-:create;autoCreateTime:milli" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

// Board is a collection of discussions.
type Board struct {
	gorm.Model
	ID          uuid.UUID     `gorm:"->;<-:create;primaryKey;unique" json:"id"`
	Name        string        `json:"name"`
	Discussions []*Discussion `json:"discussions"`
	CreatedAt   time.Time     `gorm:"->;<-:create;autoCreateTime:milli" json:"created_at"`
	UpdatedAt   time.Time     `gorm:"autoUpdateTime:milli" json:"updated_at"`
}
