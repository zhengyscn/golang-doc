package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world !!!"))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("my hello !!!"))
}

func handlerWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "my world !!!")
}

func main() {
	// 多路复用Router ServeMux
	/*
		匹配策略， 最长匹配；
	*/
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", handlerHello)
	http.HandleFunc("/world", handlerWorld)

	// http
	err := http.ListenAndServe(":8000", nil)
	log.Fatal(err)

}
