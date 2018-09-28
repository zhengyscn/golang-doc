package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world !!!"))
}

func main() {
	// 多路复用Router ServeMux
	http.HandleFunc("/", handler)

	// http
	err := http.ListenAndServe(":8000", nil)
	log.Fatal(err)

}
