package models

import (
	"coinpe/pkg/logger"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Permission struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   *time.Time     `json:"created_at,omitempty"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	Name        PermissionName `json:"name" gorm:"uniqueIndex"`
	Description string         `json:"description"`
}

type permissionRepo struct {
	db *gorm.DB
}

var (
	PermissionsToMigrate = []Permission{
		{
			ID:   1,
			Name: PermissionNameInviteUser,
		},
		{
			ID:   2,
			Name: PermissionDeactivateUser,
		},
		{
			ID:   3,
			Name: PermissionUpdateUserRole,
		},
		{
			ID:   4,
			Name: PermissionAddUserToDenyList,
		},
		{
			ID:   5,
			Name: PermissionAddFunds,
		},
		{
			ID:   6,
			Name: PermissionWriteCustomers,
		},
		{
			ID:   7,
			Name: PermissionReadInvoice,
		},
		{
			ID:   8,
			Name: PermissionWriteRole,
		},
		{
			ID:   9,
			Name: PermissionReadLedger,
		},
		{
			ID:   10,
			Name: PermissionReadPlan,
		},
		{
			ID:   11,
			Name: PermissionWritePlan,
		},
		{
			ID:   12,
			Name: PermissionReadAccount,
		},
		{
			ID:   13,
			Name: PermissionWriteAccount,
		},
	}
)

// BulkCreate implements IPermission.
func (pr *permissionRepo) BulkCreate(p []Permission) error {
	err := pr.db.
		Clauses(clause.OnConflict{DoNothing: true}).
		Model(&Permission{}).CreateInBatches(p, len(p)).Error
	if err != nil {
		logger.Error("unable to create permission ", err)
		return err
	}
	return nil
}

// Create implements IPermission.
func (pr *permissionRepo) Create(where *Permission) error {
	return pr.CreateWithTx(pr.db, where)
}

// CreateWithTx implements IPermission.
func (pr *permissionRepo) CreateWithTx(tx *gorm.DB, p *Permission) error {
	err := tx.Model(&Permission{}).Create(p).Error
	if err != nil {
		logger.Error("error in creating permission ", err)
		return err
	}
	return nil
}

// Delete implements IPermission.
func (pr *permissionRepo) Delete(where *Permission) error {
	return pr.DeleteWithTx(pr.db, where)
}

// DeleteWithTx implements IPermission.
func (pr *permissionRepo) DeleteWithTx(tx *gorm.DB, where *Permission) error {
	err := tx.Model(&Permission{}).Where(where).Delete(&Permission{}).Error
	if err != nil {
		logger.Error("unable to delete permission ", err)
		return err
	}
	return nil
}

// GetAllPermissions implements IPermission.
func (pr *permissionRepo) GetAllPermissions() ([]Permission, error) {
	var (
		permissions = []Permission{}
	)

	err := pr.db.Model(&Permission{}).
		Scan(&permissions).Error
	if err != nil {
		logger.Error("unable to get permissions | err: ", err)
		return nil, err
	}
	return permissions, nil
}

// Get implements IPermission.
func (pr *permissionRepo) Get(where *Permission) (*Permission, error) {
	return pr.GetWithTx(pr.db, where)
}

// GetWithNames implements IPermission.
func (pr *permissionRepo) GetWithNames(names []string) ([]Permission, error) {
	var (
		permissions = []Permission{}
	)

	err := pr.db.Model(&Permission{}).Where("name in (?)", names).Scan(&permissions).Error
	if err != nil {
		logger.Error("unable to get permissions with names ", err)
		return nil, err
	}
	return permissions, nil
}

// GetWithTx implements IPermission.
func (pr *permissionRepo) GetWithTx(tx *gorm.DB, where *Permission) (*Permission, error) {
	var (
		permission = Permission{}
	)
	err := tx.Model(&Permission{}).Where(where).Last(&permission).Error
	if err != nil {
		logger.Error("unable to get permission ", err)
		return nil, err
	}
	return &permission, nil
}

// Update implements IPermission.
func (pr *permissionRepo) Update(where *Permission, p *Permission) error {
	return pr.UpdateWithTx(pr.db, where, p)
}

// UpdateWithTx implements IPermission.
func (pr *permissionRepo) UpdateWithTx(tx *gorm.DB, where *Permission, p *Permission) error {
	err := tx.Model(&Permission{}).Where(where).Updates(&p).Error
	if err != nil {
		logger.Error("unable to update permission ", err)
		return nil
	}
	return err
}
