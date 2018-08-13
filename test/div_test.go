package test

import "testing"

func TestDiv(t *testing.T) {
	if ret, err := Div(1, 2); err == nil {
		if ret == 0 {
			t.Log("Sucess")
		}
	}
	if _, err := Div(1, 0); err != nil {
		t.Log(err.Error())
	}

}
