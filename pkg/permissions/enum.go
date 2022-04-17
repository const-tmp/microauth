package permissions

type Permission uint64

const (
	Access Permission = 1 << iota
	UserRead
	UserWrite
	UserFull = UserRead | UserWrite
	Admin    = Access | UserFull
)
