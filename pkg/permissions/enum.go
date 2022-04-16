package permissions

type Permission uint64

const (
	ACCESS Permission = 1 << iota
	USER_READ
	USER_WRITE
	USER_FULL = USER_READ | USER_WRITE
	ADMIN     = ACCESS | USER_FULL
)
