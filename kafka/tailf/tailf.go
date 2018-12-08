package tailf

import (
	"fmt"

	"github.com/hpcloud/tail"
)

var t *tail.Tail

func Init(filename string) (err error) {
	t, err = tail.TailFile(filename, tail.Config{Follow: true})
	if err != nil {
		panic(fmt.Sprintf("tail file err:%v", err))
	}
	return
}

func ReadLine(msgChan chan string) {
	for line := range t.Lines {
		if len(line.Text) != 0 {
			msgChan <- line.Text
		}
	}
}
