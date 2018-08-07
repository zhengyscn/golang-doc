package test

import "testing"

func TestSum(t *testing.T) {
	if Sum(1, 2) != 3 {
		t.FailNow()
	}

	if Sum(2, 3) == 5 {
		t.Log("sucess.")
	}
}
