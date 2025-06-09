package models

// Add list of model add for migrations
var migrationModels = []interface{}{
	&Account{},
	&Credential{},
	&Permission{},
	&Role{},
}

func GetMigrationModel() []interface{} {
	return migrationModels
}
