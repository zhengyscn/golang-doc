package test

import (
	"testing"
)

func TestSum(t *testing.T) {
	if Sum(1, 2) != 3 {
		t.FailNow()
	}

	if Sum(2, 3) == 5 {
		t.Log("sucess.")
	}
}

func TestSum2(t *testing.T) {
	s := []struct {
		x, y, except int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{0, 0, 0},
	}
	for _, u := range s {
		if ret := Sum(u.x, u.y); ret != u.except {
			t.Errorf("%d + %d except %d, actual result is :%d\n", u.x, u.y, u.except, ret)
		}
	}
}
