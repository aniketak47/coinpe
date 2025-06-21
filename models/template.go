package models

import (
	"time"

	"gorm.io/gorm"
)

type Template struct {
	ID        uint64         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	Name      string  `json:"name" gorm:"unique;not null"`
	ChannelID uint64  `json:"channel_id"`
	Channel   Channel `json:"-"`
	Content   string  `json:"content"`
	Subject   string  `json:"subject"`
	Category  string  `json:"category"`
}

type templateRepo struct {
	db *gorm.DB
}
