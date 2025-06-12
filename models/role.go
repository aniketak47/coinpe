package models

import (
	"coinpe/pkg/logger"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RoleID uint64

const (
	RoleSuperAdmin RoleID = iota + 1
	RoleAdmin
	RoleCustomer
)

type Role struct {
	ID        uint64         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	DisplayName   string   `json:"display_name"`
	Name          RoleType `json:"-"`
	Description   string   `json:"-"`
	SystemDefined bool     `json:"-" gorm:"default:true"`
	IsInternal    bool     `json:"-"`
	IsActive      *bool    `json:"is_active"`
	IsDefault     bool     `json:"is_default"`

	Permissions []*Permission `json:"-" gorm:"many2many:roles_permissions"`
}

type roleRepo struct {
	db *gorm.DB
}

var (
	trueVal        = true
	RolesToMigrate = []Role{
		{
			ID:            1,
			Name:          RoleTypeSuperAdmin,
			DisplayName:   "Super Admin",
			Description:   "super admin",
			SystemDefined: true,
			IsInternal:    true,
			IsActive:      &trueVal,
			Permissions: []*Permission{
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
			},
		},
		{
			ID:            2,
			Name:          RoleTypeAdmin,
			DisplayName:   "Admin",
			Description:   "admin",
			SystemDefined: true,
			IsInternal:    true,
			IsActive:      &trueVal,
			Permissions: []*Permission{
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
			},
		},
		{
			ID:            3,
			Name:          RoleTypeCustomer,
			DisplayName:   "Customer",
			Description:   "customer",
			SystemDefined: true,
			IsActive:      &trueVal,
		},
	}
)

// GetAllExtenralRoles implements IRole.
func (r *roleRepo) GetAllExternalRoles() ([]Role, error) {
	var (
		roles = []Role{}
	)

	err := r.db.Model(&Role{}).
		Where("is_internal = ?", false).
		Scan(&roles).Error
	if err != nil {
		logger.Error("unable to get roles ", err)
		return nil, err
	}
	return roles, nil
}

// CheckIfPermissionExists implements IRole.
func (r *roleRepo) CheckIfPermissionExists(roleID uint64, permissionName PermissionName) (bool, error) {
	var (
		role = Role{}
	)
	err := r.db.Model(&Role{}).
		Preload("Permissions", func(tx *gorm.DB) *gorm.DB {
			return tx.Where("name = ?", permissionName)
		}).Where(&Role{
		ID: uint64(roleID),
	}).Last(&role).Error
	if err != nil {
		logger.Error("unable to check if permission exists ", err)
		return false, nil
	}
	if len(role.Permissions) == 0 {
		return false, nil
	}
	return true, nil

}

func (r *roleRepo) GetByID(ID uint64) (*Role, error) {
	return r.GetWithTx(r.db, &Role{ID: ID})
}

func (r *roleRepo) First(where *Role) (*Role, error) {
	return r.GetWithTx(r.db, where)
}

func (r *roleRepo) Find(where *Role) (*[]Role, error) {
	var m []Role
	err := r.db.Model(&Role{}).Where(where).Find(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *roleRepo) GetWithTx(tx *gorm.DB, where *Role) (*Role, error) {
	var m Role
	err := tx.Preload("Permissions").Model(&Role{}).Where(where).First(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *roleRepo) Create(u *Role) error {
	return r.CreateWithTx(r.db, u)
}

func (r *roleRepo) CreateWithTx(tx *gorm.DB, u *Role) error {
	err := tx.Create(u).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *roleRepo) BulkCreate(roles *[]Role) error {
	err := r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(roles).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *roleRepo) Update(u *Role, ID uint64) error {
	return r.UpdateWithTx(r.db, u, ID)
}

func (r *roleRepo) UpdateWithTx(tx *gorm.DB, u *Role, ID uint64) error {
	err := tx.Model(&Role{ID: ID}).Updates(u).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *roleRepo) Delete(ID uint64) error {
	err := r.db.Where(&Role{ID: ID}).Delete(&Role{}).Error
	if err != nil {
		return err
	}
	return nil
}
