package balance

import (
	"errors"
)

type RoundrobinBalance struct {
	currIndex int
}

func (this *RoundrobinBalance) DoBalance(instances []*Instnace) (*Instnace, error) {
	var length = len(instances)
	if length == 0 {
		return nil, errors.New("No instnaces")
	}

	instance := instances[this.currIndex]
	this.currIndex = (this.currIndex + 1) % length
	return instance, nil
}
