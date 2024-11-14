package pkg

import "testing"

func TestEncryptPassword(t *testing.T) {
	password := EncryptPassword("123adsd")
	t.Logf("your password : %v\n", password)
}
