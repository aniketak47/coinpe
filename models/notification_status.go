package models

import (
	"coinpe/pkg/logger"
	"coinpe/pkg/utils"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NotificationStatus struct {
	ID        uint64         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	Code        uint   `json:"code" gorm:"unique"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type notificationStatusRepo struct {
	db *gorm.DB
}

var (
	// Notification Status Code
	NotificationStatusQueuedCode    = uint(100)
	NotificationStatusSentCode      = uint(101)
	NotificationStatusDeliveredCode = uint(102)
	NotificationStatusReadCode      = uint(103)
	NotificationStatusFailedCode    = uint(104)

	// Notification Status Description
	NotificationStatusQueuedDescription    = "queued the mail"
	NotificationStatusSentDescription      = "sent the mail"
	NotificationStatusDeliveredDescription = "delivered the mail"
	NotificationStatusReadDescription      = "read the mail"
	NotificationStatusFailedDescription    = "failed to deliver"

	// Notification Status Name
	NotificationStatusQueuedName    = "Queued"
	NotificationStatusSentName      = "Sent"
	NotificationStatusDeliveredName = "Delivered"
	NotificationStatusReadName      = "Read"
	NotificationStatusFailedName    = "Failed"

	NotificationStatuses = []NotificationStatus{
		{
			ID:          1,
			CreatedAt:   utils.TimePTR(time.Now()),
			UpdatedAt:   utils.TimePTR(time.Now()),
			Code:        NotificationStatusQueuedCode,
			Description: NotificationStatusQueuedDescription,
			Name:        NotificationStatusQueuedName,
		},
		{
			ID:          2,
			CreatedAt:   utils.TimePTR(time.Now()),
			UpdatedAt:   utils.TimePTR(time.Now()),
			Code:        NotificationStatusSentCode,
			Description: NotificationStatusSentDescription,
			Name:        NotificationStatusSentName,
		},
		{
			ID:          3,
			CreatedAt:   utils.TimePTR(time.Now()),
			UpdatedAt:   utils.TimePTR(time.Now()),
			Code:        NotificationStatusDeliveredCode,
			Description: NotificationStatusDeliveredDescription,
			Name:        NotificationStatusDeliveredName,
		},
		{
			ID:          4,
			CreatedAt:   utils.TimePTR(time.Now()),
			UpdatedAt:   utils.TimePTR(time.Now()),
			Code:        NotificationStatusReadCode,
			Description: NotificationStatusReadDescription,
			Name:        NotificationStatusReadName,
		},
		{
			ID:          5,
			CreatedAt:   utils.TimePTR(time.Now()),
			UpdatedAt:   utils.TimePTR(time.Now()),
			Code:        NotificationStatusFailedCode,
			Description: NotificationStatusFailedDescription,
			Name:        NotificationStatusFailedName,
		},
	}
)

func (r *notificationStatusRepo) BulkCreate(nStatus []NotificationStatus) error {
	err := r.db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(nStatus, len(nStatus)).Error
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
