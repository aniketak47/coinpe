package models

import (
	"coinpe/pkg/logger"
	"coinpe/pkg/utils"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Account struct {
	ID        uint           `json:"id,omitempty" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitzero" gorm:"index"`

	UUID        string  `json:"uuid,omitempty" gorm:"unique"`
	FirstName   string  `json:"first_name,omitempty"`
	LastName    string  `json:"last_name,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty" gorm:"unique"`
	Email       string  `json:"email,omitempty"`

	RoleID uint64 `json:"role_id" gorm:"not null;index"`
	Role   *Role  `json:"role,omitempty"`

	Credentials []Credential `json:"-"`
}

type accountRepo struct {
	db *gorm.DB
}

func (a *Account) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.AddClause(clause.OnConflict{
		DoNothing: true,
	})

	userUUID, err := utils.GenerateNanoID(12, "acc_")
	if err != nil {
		logger.Error("unable to generate nano id | err: ", err)
		return err
	}

	if a.UUID != "" {
		userUUID = a.UUID
	}

	a.UUID = userUUID

	totpAccountName := userUUID
	if a.PhoneNumber != nil && *a.PhoneNumber != "" {
		totpAccountName = *a.PhoneNumber
	}

	ac, err := CreateOTPSecret(totpAccountName)
	if err != nil {
		return err
	}
	a.Credentials = append(a.Credentials, *ac)

	return nil
}

func (ar *accountRepo) Get(where *Account) (*Account, error) {
	return ar.GetWithTx(ar.db, where)
}

// GetWithTx implements IAccount.
func (ar *accountRepo) GetWithTx(tx *gorm.DB, where *Account) (*Account, error) {
	var (
		u = Account{}
	)
	err := tx.Model(&Account{}).
		Where(where).
		Last(&u).Error
	if err != nil {
		logger.Error("unable to query account | err: ", err)
		return nil, err
	}
	return &u, nil
}

// Create implements IAcount.
func (ar *accountRepo) Create(u *Account) error {
	return ar.CreateWithTx(ar.db, u)
}

// CreateWithTx implements IAccount.
func (ar *accountRepo) CreateWithTx(tx *gorm.DB, u *Account) error {
	err := tx.Model(&Account{}).Create(&u).Error
	if err != nil {
		logger.Error("unable to create account | err: ", err)
		return err
	}
	return nil
}

func (ar *accountRepo) Delete(userID uint) error {
	return ar.DeleteWithTx(ar.db, &Account{ID: userID})
}

func (ar *accountRepo) DeleteWithTx(tx *gorm.DB, u *Account) error {
	err := tx.Model(&Account{}).
		Where(u).
		Delete(&Account{}).
		Error
	if err != nil {
		logger.Error("unable to delete account | err: ", err)
		return err
	}
	return nil
}

// Update implements IAccount.
func (ar *accountRepo) Update(where *Account, a *Account) error {
	return ar.UpdateWithTx(ar.db, where, a)
}

// UpdateWithTx implements IAccount.
func (ar *accountRepo) UpdateWithTx(tx *gorm.DB, where *Account, a *Account) error {
	err := tx.Model(&Account{}).
		Where(where).Updates(&a).Error
	if err != nil {
		logger.Error("unable to update account | err: ", err)
		return err
	}
	return nil
}

// FindOne implements IAccount.
func (ar *accountRepo) FindOne(tx *gorm.DB, email, phoneNumber, accountUUID string) (*Account, error) {
	var (
		account = Account{}
	)

	builder := tx.Model(&Account{}).Preload("Credentials")

	if phoneNumber != "" {
		builder.Or(&Account{
			PhoneNumber: &phoneNumber,
		})
	}

	if accountUUID != "" {
		builder.Or(&Account{
			UUID: accountUUID,
		})
	}

	if email != "" {
		builder.Or(&Account{
			Email: email,
		})
	}

	err := builder.Find(&account).Error
	if err != nil {
		logger.Error("unable to find account | err: ", err)
		return nil, err
	}

	return &account, nil
}

func (ar *accountRepo) GetWithCredentials(where *Account, credcredentialType CredentialsTypeSlug) (*Account, error) {
	var (
		u = Account{}
	)
	err := ar.db.Model(&Account{}).
		Preload("Credentials", "Type = ?", credcredentialType).
		Where(where).
		Last(&u).Error
	if err != nil {
		logger.Error("unable to query account | err: ", err)
		return nil, err
	}
	return &u, nil
}
