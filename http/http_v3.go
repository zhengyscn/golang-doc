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

type myHello struct {
}

func (h *myHello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("my hello !!!"))
}

type myWorld struct {
}

func (h *myWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("my world !!!"))
}

func main() {
	handler := &myHandle{}
	handlerHello := &myHello{}
	handlerWorld := &myWorld{}
	s := &http.Server{
		Addr: ":8000",
		//Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.Handle("/", handler)
	http.Handle("/hello", handlerHello)
	http.Handle("/world", handlerWorld)
	log.Fatal(s.ListenAndServe())
}
