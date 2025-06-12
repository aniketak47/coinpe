package models

import "gorm.io/gorm"

func InitAccountRepo(db *gorm.DB) IAccount {
	return &accountRepo{
		db: db,
	}
}

func InitCredentialRepo(db *gorm.DB) ICredential {
	return &credentialRepo{
		db: db,
	}
}

func InitPermissionRepo(db *gorm.DB) IPermission {
	return &permissionRepo{
		db: db,
	}
}

func InitRoleRepo(db *gorm.DB) IRole {
	return &roleRepo{
		db: db,
	}
}

func InitWalletRepo(DB *gorm.DB) IWallet {
	return &walletRepo{
		db: DB,
	}
}
