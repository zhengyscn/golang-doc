package balance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
}

func (this *RandomBalance) DoBalance(instances []*Instnace) (*Instnace, error) {
	var length = len(instances)
	if length == 0 {
		return nil, errors.New("No instnaces")
	}
	instance := instances[rand.Intn(length)]
	return instance, nil
}
