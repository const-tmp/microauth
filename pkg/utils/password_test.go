package utils

import "testing"

func TestPassword(t *testing.T) {
	pw := "testtt"
	hash, err := HashPassword(pw)
	if err != nil {
		t.Error(err)
	}
	t.Log(hash, len(hash))
	ok := CheckPasswordHash(pw, hash)
	t.Log(ok)
}
