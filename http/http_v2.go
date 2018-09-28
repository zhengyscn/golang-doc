package main

import (
	"log"
	"net/http"
	"time"
)

type myHandle struct {
}

func (h *myHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world !!!"))
}

func main() {
	handler := &myHandle{}
	s := &http.Server{
		Addr:           ":8000",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
