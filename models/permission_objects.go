package models

type PermissionName string

const (
	PermissionNameInviteUser    PermissionName = "INVITE_USER"
	PermissionDeactivateUser    PermissionName = "DEACTIVATE_USER"
	PermissionUpdateUserRole    PermissionName = "UPDATE_USER_ROLE"
	PermissionAddUserToDenyList PermissionName = "ADD_USER_TO_DENY_LIST"
	PermissionAddFunds          PermissionName = "ADD_FUNDS"
	PermissionWriteCustomers    PermissionName = "WRITE_CUSTOMERS"
	PermissionReadInvoice       PermissionName = "READ_INVOICE"
	PermissionWriteRole         PermissionName = "WRITE_ROLE"
	PermissionReadLedger        PermissionName = "READ_LEDGER"
	PermissionReadPlan          PermissionName = "READ_PLAN"
	PermissionWritePlan         PermissionName = "WRITE_PLAN"
	PermissionReadAccount       PermissionName = "READ_ACCOUNT"
	PermissionWriteAccount      PermissionName = "WRITE_ACCOUNT"
)
