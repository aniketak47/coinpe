package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Provider struct {
	ID        uint64         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitzero" gorm:"index"`

	Name              string         `json:"name"`
	Key               string         `json:"key"`
	TemplateSupported bool           `json:"template_supported"`
	ChannelID         uint64         `json:"channel_id"`
	Channel           Channel        `json:"-"`
	Weight            uint64         `json:"weight"`
	Config            datatypes.JSON `json:"config"`
}

type providerRepo struct {
	db *gorm.DB
}
