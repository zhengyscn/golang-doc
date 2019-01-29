package main

import (
	"fmt"

	"github.com/satori/go.uuid"
)

func genUUID() (string, error) {
	u2, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("Something went wrong: %v", err)
	}
	return u2.String(), nil
}

func main() {
	uid, err := genUUID()
	if err != nil {
		fmt.Printf("gen uuid err:%v\n", err)
		return
	}
	fmt.Printf("gen uuid succ, uuid : %s\n", uid)
}
