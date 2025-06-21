package models

import (
	"time"

	"gorm.io/gorm"
)

type NotificationUpdate struct {
	ID        uint64         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitzero" gorm:"index"`

	StatusCode         uint               `json:"status_code"`
	Description        string             `json:"description"`
	NotificationStatus NotificationStatus `json:"-" gorm:"foreignKey:StatusCode;references:Code"`
	NotificationID     uint64             `json:"notifications_id"`
	Notification       Notification       `json:"-"`
}

type notificationUpdateRepo struct {
	db *gorm.DB
}

func (p *notificationUpdateRepo) GetNotificationUpdateByNotificationID(tx *gorm.DB, NotificationID uint64) (*[]NotificationUpdate, error) {
	var o []NotificationUpdate
	err := tx.Model(&NotificationUpdate{}).Where(NotificationUpdate{NotificationID: NotificationID}).Find(&o).Error
	if err != nil {
		return nil, err
	}
	return &o, nil
}
func (p *notificationUpdateRepo) GetNotificationUpdateWithTx(tx *gorm.DB, where *NotificationUpdate) (*NotificationUpdate, error) {
	var o NotificationUpdate
	err := tx.Model(&NotificationUpdate{}).Where(where).First(&o).Error
	if err != nil {
		return nil, err
	}
	return &o, nil
}

func (p *notificationUpdateRepo) CreateNotificationUpdate(u *NotificationUpdate) error {
	return p.CreateNotificationUpdateWithTx(p.db, u)
}
func (p *notificationUpdateRepo) CreateNotificationUpdateWithTx(tx *gorm.DB, u *NotificationUpdate) error {
	err := tx.Create(u).Error
	if err != nil {
		return err
	}
	return nil
}
