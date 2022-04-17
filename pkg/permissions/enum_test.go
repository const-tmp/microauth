package permissions

import "testing"

func TestEnum(t *testing.T) {
	t.Logf("%064b\n", Access)
	t.Logf("%064b\n", UserRead)
	t.Logf("%064b\n", UserWrite)
	t.Logf("%064b\n", UserFull)
	t.Logf("%064b\n", Admin)
}
