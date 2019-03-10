package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func sayhelloName(write http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println(request.Form)
	fmt.Println("path", request.URL.Path)
	fmt.Println("scheme", request.URL.Scheme)
	fmt.Println(request.Form["url_long"])

	for k, v := range request.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(write, "Hello astaxie!")
}

func viewer(write http.ResponseWriter, request *http.Request) {
	fmt.Println("method:", request.Method)
	if request.Method == "GET" {
		t, err := template.ParseFiles("../../viewer.html")
		if err != nil {
			http.Error(write, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(write, nil)
	} else {
		fmt.Println("username:", request.PostForm["username"])
	}
}

func getUser(write http.ResponseWriter, request *http.Request) {
	fmt.Println("method:", request.Method)
	if request.Method == "GET" {
		t, err := template.ParseFiles("../../viewer.html")
		if err != nil {
			http.Error(write, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(write, nil)
	} else {
		request.ParseForm()
		fmt.Println("username:", request.FormValue("username"))
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9000",
	}

	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/viewer", viewer)
	http.HandleFunc("/getUser", getUser)
	server.ListenAndServe()
}
