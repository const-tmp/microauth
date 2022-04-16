package permissions

import "testing"

func TestEnum(t *testing.T) {
	t.Logf("%064b\n", ACCESS)
	t.Logf("%064b\n", USER_READ)
	t.Logf("%064b\n", USER_WRITE)
	t.Logf("%064b\n", USER_FULL)
	t.Logf("%064b\n", ADMIN)
}
