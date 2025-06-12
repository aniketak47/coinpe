package models

import (
	"coinpe/pkg/constants"

	"gorm.io/gorm"
)

// AddSystemData: Use this hook to populate any default data to the database.
func AddSystemData(db *gorm.DB, env constants.AppEnv) {
	InitPermissionRepo(db).BulkCreate(PermissionsToMigrate)
	InitRoleRepo(db).BulkCreate(&RolesToMigrate)
	InitWalletRepo(db).Create(&CoinpeWallet)
}
