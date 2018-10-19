package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	host = "http://localhost:8080"
	hostHttps = "http://localhost:8080"
)

func reqGet() {
	path := "/html"
	url := fmt.Sprintf("%s%s", host, path)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if resp.StatusCode == http.StatusOK {
		if bytes, err := ioutil.ReadAll(resp.Body); err == nil {
			fmt.Println(string(bytes))
		}

	}

}

func reqGetV2() {
	path := "/html"
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", host, path), nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	req.Header.Add("If-None-Match", "/")
	resp, err := client.Do(req)
	if resp.StatusCode == http.StatusOK {
		if bytes, err := ioutil.ReadAll(resp.Body); err == nil {
			fmt.Println(string(bytes))
		}
		
	}
	
}

func reqGetHttps() {
	tr := &http.Transport{
		TLSClientConfig:&tls.Config{InsecureSkipVerify:true}
	}
	client := http.Client{
		Transport:tr,
	}
	path := "/html"
	resp, err := client.Get(fmt.Sprintf("%s%s", hostHttps, path))
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if resp.StatusCode == http.StatusOK {
		if bytes, err := ioutil.ReadAll(resp.Body); err == nil {
			fmt.Println(string(bytes))
		}
		
	}
	
}

func reqPost() {
	path := "/form/v2"
	resp, err := http.PostForm(fmt.Sprintf("%s%s", host, path), url.Values{
		"username": {"zhengyscn"},
		"password": {"123456"},
	})
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if resp.StatusCode == http.StatusOK {
		if bytes, err := ioutil.ReadAll(resp.Body); err == nil {
			fmt.Println(string(bytes))
		}

	}
}

func main() {
	reqGet()
	fmt.Println("-------------")
	reqPost()
}
