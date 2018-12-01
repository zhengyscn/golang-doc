package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/context"
)

type RespData struct {
	err  error
	resp *http.Response
}

func doCall(ctx context.Context) {
	var wg sync.WaitGroup

	transport := http.Transport{}
	client := http.Client{
		Transport: &transport,
	}

	respChan := make(chan *RespData)
	req, err := http.NewRequest("GET", "http://127.0.0.1:8000", nil)
	if err != nil {
		fmt.Printf("new request failed, err:%v\n", err)
		return
	}
	wg.Add(1)
	defer wg.Wait()
	go func() {
		defer wg.Done()
		resp, err := client.Do(req)
		respData := &RespData{
			err:  err,
			resp: resp,
		}
		respChan <- respData

	}()
DONE:
	select {
	case <-ctx.Done(): // timeout
		transport.CancelRequest(req)
		fmt.Println("do call api timeout")
		break DONE
	case result := <-respChan:
		fmt.Printf("call server api succ\n")
		if result.err != nil {
			fmt.Printf("call api failed, err:%v\n", err)
			return
		}
		data, _ := ioutil.ReadAll(result.resp.Body)
		fmt.Printf("resp:%v\n", string(data))
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()
	for {
		doCall(ctx)
	}

}
