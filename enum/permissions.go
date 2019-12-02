package enum

const (
	RoleGuest = iota + 1
	RoleDeveloper
	RoleSRE
	RolePM
	RoleOwner
	RoleAdmin
)

const (
	PermRead = iota + 1
	PermWrite
	PermCreate
	PermDelete
	PermUpdate
)
