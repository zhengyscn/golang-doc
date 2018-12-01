package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"log"
)

/*
	慢请求
*/

func handle(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(2)
	if number == 0 {
		time.Sleep(time.Second * 10)
		fmt.Fprintf(w, "slow quest.\n")
		return
	}
	fmt.Fprint(w, "quick quest.\n")

}

func main() {
	http.HandleFunc("/", handle)
	err := http.ListenAndServe(":8000", nil)
	log.Fatal(err)
}
