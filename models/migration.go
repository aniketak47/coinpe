package models

// Add list of model add for migrations
var migrationModels = []interface{}{}

func GetMigrationModel() []interface{} {
	return migrationModels
}
