package models

import (
	"coinpe/pkg/logger"
	"coinpe/pkg/utils"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Channel struct {
	ID        uint64         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitzero" gorm:"index"`

	Name string `json:"name" gorm:"unique"`
}

type channelRepo struct {
	db *gorm.DB
}

var (
	c1 = Channel{
		ID:   1,
		Name: "email",
	}
	c2 = Channel{
		ID:   2,
		Name: "sms",
	}
	c3 = Channel{
		ID:   3,
		Name: "push",
	}
	c4 = Channel{
		ID:   4,
		Name: "whatsapp",
	}

	ChannelList = []Channel{c1, c2, c3, c4}
)

// change channel id to channel name
var (
	ChannelEmailName    = "email"
	ChannelSmsName      = "sms"
	ChannelPushName     = "push"
	ChannelWhatsappName = "whatsapp"

	Channels = []Channel{
		{
			ID:        1,
			CreatedAt: utils.TimePTR(time.Now()),
			UpdatedAt: utils.TimePTR(time.Now()),
			Name:      ChannelEmailName,
		},
		{
			ID:        2,
			CreatedAt: utils.TimePTR(time.Now()),
			UpdatedAt: utils.TimePTR(time.Now()),
			Name:      ChannelSmsName,
		},
		{
			ID:        3,
			CreatedAt: utils.TimePTR(time.Now()),
			UpdatedAt: utils.TimePTR(time.Now()),
			Name:      ChannelPushName,
		},
		{
			ID:        4,
			CreatedAt: utils.TimePTR(time.Now()),
			UpdatedAt: utils.TimePTR(time.Now()),
			Name:      ChannelWhatsappName,
		},
	}
	ChannelNameToIDMap = map[string]uint64{}
	ChannelIDToNameMap = map[uint64]string{}
)

func init() {
	for _, val := range Channels {
		ChannelNameToIDMap[val.Name] = val.ID
		ChannelIDToNameMap[val.ID] = val.Name
	}
}

func (c *channelRepo) GetByID(ID uint64) (*Channel, error) {
	return c.GetByIDWithTx(c.db, &Channel{ID: ID})
}

func (c *channelRepo) GetByIDWithTx(tx *gorm.DB, where *Channel) (*Channel, error) {
	var o Channel
	err := tx.Model(&Channel{}).Where(where).First(&o).Error
	if err != nil {
		return nil, err
	}
	return &o, nil

}

func (p *channelRepo) BulkCreate(channels []Channel) error {
	err := p.db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(channels, len(channels)).Error
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (c *channelRepo) Create(u *Channel) error {
	return c.CreateWithTx(c.db, u)
}

func (c *channelRepo) CreateWithTx(tx *gorm.DB, u *Channel) error {
	err := tx.Create(u).Error
	if err != nil {
		return err
	}
	return nil
}
