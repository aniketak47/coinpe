package models

import (
	"gorm.io/gorm"
)

type IAccount interface {
	Get(where *Account) (*Account, error)
	GetWithTx(tx *gorm.DB, where *Account) (*Account, error)
	Create(u *Account) error
	CreateWithTx(tx *gorm.DB, u *Account) error
	Update(where *Account, a *Account) error
	UpdateWithTx(tx *gorm.DB, where *Account, a *Account) error
	Delete(userID uint) error
	DeleteWithTx(tx *gorm.DB, u *Account) error
	FindOne(tx *gorm.DB, email, phoneNumber, accountUUID string) (*Account, error)
	GetWithCredentials(where *Account, credentialType CredentialsTypeSlug) (*Account, error)
}

type ICredential interface {
	Create(c *Credential) error
	CreateWithTx(tx *gorm.DB, c *Credential) error
	Get(where *Credential) (*Credential, error)
	GetWithTx(tx *gorm.DB, where *Credential) (*Credential, error)
	Update(where *Credential, u *Credential) error
	UpdateWithTx(tx *gorm.DB, where *Credential, u *Credential) error
	Delete(where *Credential) error
	DeleteWithTx(tx *gorm.DB, where *Credential) error
	CheckIfPasswordIsValid(userID uint, password string) (bool, error)
}

type IPermission interface {
	Create(where *Permission) error
	BulkCreate(p []Permission) error
	CreateWithTx(tx *gorm.DB, p *Permission) error
	GetWithTx(tx *gorm.DB, where *Permission) (*Permission, error)
	Get(where *Permission) (*Permission, error)
	UpdateWithTx(tx *gorm.DB, where *Permission, p *Permission) error
	Update(where *Permission, p *Permission) error
	Delete(where *Permission) error
	DeleteWithTx(tx *gorm.DB, where *Permission) error
	GetWithNames(names []string) ([]Permission, error)
	GetAllPermissions() ([]Permission, error)
}

type IRole interface {
	GetByID(ID uint64) (*Role, error)
	First(where *Role) (*Role, error)
	Find(where *Role) (*[]Role, error)
	GetWithTx(tx *gorm.DB, where *Role) (*Role, error)
	Create(u *Role) error
	CreateWithTx(tx *gorm.DB, u *Role) error
	BulkCreate(roles *[]Role) error
	Update(u *Role, ID uint64) error
	UpdateWithTx(tx *gorm.DB, u *Role, ID uint64) error
	Delete(ID uint64) error
	CheckIfPermissionExists(roleID uint64, permissionName PermissionName) (bool, error)
	GetAllExternalRoles() ([]Role, error)
}

type IWallet interface {
	Get(where *Wallet) (*Wallet, error)
	GetWithTx(tx *gorm.DB, where *Wallet) (*Wallet, error)
	Create(w *Wallet) error
	CreateWithTx(tx *gorm.DB, w *Wallet) error
	Delete(walletID string) error
	UpdateBalance(tx *gorm.DB, walletID string, totalBalance int) error
	CanDebit(w *Wallet, amountInCents int) bool
	UpdateWithTx(tx *gorm.DB, where *Wallet, w *Wallet) error
	Update(where *Wallet, w *Wallet) error
}
