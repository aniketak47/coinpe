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

func InitWalletRepo(db *gorm.DB) IWallet {
	return &walletRepo{
		db: db,
	}
}

func InitChannelRepo(db *gorm.DB) IChannel {
	return &channelRepo{
		db: db,
	}
}

func InitProviderRepo(db *gorm.DB) IProvider {
	return &providerRepo{
		db: db,
	}
}

func InitTemplateRepo(db *gorm.DB) ITemplate {
	return &templateRepo{
		db: db,
	}
}

func InitNotificationStatusRepo(DB *gorm.DB) INotificationStatus {
	return &notificationStatusRepo{
		db: DB,
	}
}

func InitNotificationRepo(DB *gorm.DB) INotification {
	return &notificationRepo{
		db: DB,
	}
}

func InitNotificationUpdateRepo(DB *gorm.DB) INotificationUpdate {
	return &notificationUpdateRepo{
		db: DB,
	}
}
