package main

import (
	"html/template"
	"net/http"
)

type Person struct {
	UserName string
	Email    string
}

func tmpl(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("testTemplate.gtpl")
	if err != nil {
		panic(err)
	}

	p := Person{UserName: "Astaxie", Email: "123@456.com"}
	t1.Execute(w, p)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/testTemplate", tmpl)
	server.ListenAndServe()
}
