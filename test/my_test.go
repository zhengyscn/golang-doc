package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	fmt.Println("setup")
	rand.Seed(time.Now().Unix())
	m.Run()
	fmt.Println("tear down")
}
