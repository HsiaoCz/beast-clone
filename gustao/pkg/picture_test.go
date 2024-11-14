package pkg

import "testing"

func TestGetPicture(t *testing.T) {
	picture := GetPicture("ATR")
	t.Logf("your picture : %v\n", picture)
}
