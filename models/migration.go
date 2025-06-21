package models

// Add list of model add for migrations
var migrationModels = []interface{}{
	&Account{},
	&Credential{},
	&Permission{},
	&Role{},
	&Wallet{},
	&Channel{},
	&Provider{},
	&Template{},
	&NotificationStatus{},
	&Notification{},
	&NotificationUpdate{},
}

func GetMigrationModel() []interface{} {
	return migrationModels
}
