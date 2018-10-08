package main

import (
	"fmt"
	"io"
	"net/http"
)

func bodyRecv(w http.ResponseWriter, r *http.Request) {
	length := r.ContentLength
	buf := make([]byte, length)
	if _, err := r.Body.Read(buf); err == nil || err == io.EOF {
		fmt.Fprintf(w, string(buf))
	} else {
		fmt.Fprintf(w, "error")
	}

}

func handForm_v1(w http.ResponseWriter, r *http.Request) {
	/*
		curl -id 'username=zhengyscn&password=123456' http://127.0.0.1:8080/form?remember=true
	*/
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.PostForm)

	fmt.Fprintf(w, "ok")

}

func handForm_v2(w http.ResponseWriter, r *http.Request) {
	/*
		curl -id 'username=zhengyscn&password=123456' http://127.0.0.1:8080/form?remember=true
	*/
	fmt.Println(r.FormValue("username"))
	fmt.Println(r.FormValue("password"))
	fmt.Println("-----")
	fmt.Printf("%#v\n", r.PostForm)
	fmt.Fprintf(w, "ok")
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormFile("fd"))
	fmt.Fprintf(w, "upload file ok.")
}

func main() {
	http.HandleFunc("/", bodyRecv)
	http.HandleFunc("/form/v1", handForm_v1)
	http.HandleFunc("/form/v2", handForm_v2)
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}
