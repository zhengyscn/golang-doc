package test

import "fmt"

func Div(i, j int) (int, error) {
	if j == 0 {
		return 0, fmt.Errorf("j is zero")
	}
	return i / j, nil
}
