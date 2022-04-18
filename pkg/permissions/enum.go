package permissions

type Permissions uint64

const (
	Access Permissions = 1 << iota
	UserRead
	UserWrite
	UserFull = UserRead | UserWrite
	Admin    = Access | UserFull
)

func (p Permissions) Check(userPermissions Permissions) bool {
	if userPermissions&p == p {
		return true
	}
	return false
}
