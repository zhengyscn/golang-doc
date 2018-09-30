package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
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

/*
	装饰器
*/
func logHTTP(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("handler func called -", name)
		h(writer, request)
	}
}

func main() {
	// 多路复用Router ServeMux
	/*
		匹配策略， 最长匹配；
	*/
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", handlerHello)
	http.HandleFunc("/world", logHTTP(handlerWorld))

	// http
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)

}
