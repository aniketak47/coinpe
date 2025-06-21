package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID        uint64         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitzero" gorm:"index"`

	Content             string     `json:"content"`
	ProviderID          uint64     `json:"provider_id"`
	Provider            Provider   `json:"-"`
	FromEmail           string     `json:"from_email,omitempty"`
	ToEmail             string     `json:"to_email,omitempty"`
	FromSid             string     `json:"from_sid"`
	UniqueRequestID     string     `gorm:"index" json:"unique_request_id"`
	ToMobileNumber      string     `json:"to_mobile_number,omitempty"`
	CurrentStatusCode   uint64     `json:"current_status_code"`
	CurrentStatus       string     `json:"current_status"`
	TemplateID          uint64     `json:"template_id"`
	ProviderReferenceID string     `json:"provider_reference_id"`
	DispatchedAt        *time.Time `json:"dispatched_at"`
	DeliveredAt         *time.Time `json:"delivered_at"`
	ReadAt              *time.Time `json:"read_at"`
	Subject             string     `json:"subject"`
}

type notificationRepo struct {
	db *gorm.DB
}

func (p *notificationRepo) GetNotificationByTemplateID(TemplateID uint64) (*Notification, error) {
	return p.GetNotificationWithTx(p.db, &Notification{TemplateID: TemplateID})
}
func (p *notificationRepo) GetNotificationWithTx(tx *gorm.DB, where *Notification) (*Notification, error) {
	var o Notification
	err := tx.Model(&Notification{}).Where(where).First(&o).Error
	if err != nil {
		return nil, err
	}
	return &o, nil
}

func (p *notificationRepo) CreateNotification(u *Notification) error {
	return p.CreateNotificationWithTx(p.db, u)
}
func (p *notificationRepo) CreateNotificationWithTx(tx *gorm.DB, u *Notification) error {
	err := tx.Create(u).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *notificationRepo) Update(n *Notification, ID uint64) error {
	return p.UpdateWithTx(p.db, n)
}
func (p *notificationRepo) UpdateWithTx(tx *gorm.DB, n *Notification) error {
	err := tx.Model(&Notification{ID: n.ID}).Updates(n).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *notificationRepo) GetByProviderReferenceID(providerReferenceID string) (*Notification, error) {
	var o Notification
	err := p.db.Model(&Notification{}).Where(&Notification{ProviderReferenceID: providerReferenceID}).First(&o).Error
	if err != nil {
		return &o, err
	}
	return &o, nil
}
