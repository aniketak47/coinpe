package models

// Add list of model add for migrations
var migrationModels = []interface{}{
	&Account{},
	&Credential{},
	&Permission{},
	&Role{},
	&Wallet{},
}

func GetMigrationModel() []interface{} {
	return migrationModels
}
