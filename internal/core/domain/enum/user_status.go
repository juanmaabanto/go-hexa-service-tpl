package enum

type UserStatus int

const (
	Inactive UserStatus = iota
	Active
	PasswordChangeRequired
)
