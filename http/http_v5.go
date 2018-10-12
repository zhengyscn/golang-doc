package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	fmt.Println("-----")
	fmt.Println(r.PostFormValue("username"))
	fmt.Println(r.PostFormValue("password"))
	fmt.Fprintf(w, "ok")
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	/*
		curl -i --form upload=@/etc/passwd http://127.0.0.1:8080/upload
	*/
	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["upload"][0]
	if fd, err := fileHeader.Open(); err == nil {
		bts, _ := ioutil.ReadAll(fd)
		fmt.Fprintf(w, string(bts))
	} else {
		fmt.Fprintf(w, err.Error())
	}
}

func jsonRecv(w http.ResponseWriter, r *http.Request) {
	/*
		curl -i -H "Content-Type: application/json" -d '{"username" : "x1", "password" : "123"}'  http://127.0.0.1:8080/json
	*/
	type jsonData struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Age      int    `json:"age,string"`
	}

	var data jsonData

	if bytes, err := ioutil.ReadAll(r.Body); err == nil {
		if err = json.Unmarshal(bytes, &data); err == nil {
			fmt.Printf("%#v\n", data)
		} else {
			fmt.Printf("Json unmarshal error %v\n", err.Error())
		}
	} else {
		fmt.Printf("ReadAll error %v\n", err.Error())
	}
	fmt.Fprintf(w, "ok")

}

func readHtml(w http.ResponseWriter, r *http.Request) {
	if bytes, err := ioutil.ReadFile("templates/index.html"); err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request."))
	}
}

func statusMonitor(w http.ResponseWriter, r *http.Request) {
	/*
		Todo
	*/
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("no such service."))
}

func httpRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "/html")
	w.WriteHeader(http.StatusFound)
}

func main() {
	http.HandleFunc("/", bodyRecv)
	http.HandleFunc("/form/v1", handForm_v1)
	http.HandleFunc("/form/v2", handForm_v2)
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/json", jsonRecv)
	http.HandleFunc("/html", readHtml)
	http.HandleFunc("/status", statusMonitor)
	http.HandleFunc("/redirect", httpRedirect)
	http.ListenAndServe(":8080", nil)
}
