package models

type RoleType string

const (
	RoleTypeAdmin      RoleType = "ADMIN"
	RoleTypeCustomer   RoleType = "CUSTOMER"
	RoleTypeSuperAdmin RoleType = "SUPER_ADMIN"
)
