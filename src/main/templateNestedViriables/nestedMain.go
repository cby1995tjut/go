package main

import (
	"html/template"
	"net/http"
)

func tmpl(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("nested.gtpl")
	if err != nil {
		panic(err)
	}

	task := Task{1, "title test", "hello content", "2022"}
	context := Context{Tasks: []Task{task}, Navigation: "hell navi"}
	t1.Execute(w, context)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/nested", tmpl)
	server.ListenAndServe()
}

type Task struct {
	Id      int
	Title   string
	Content string
	Created string
}

//Context is the struct passed to templates
type Context struct {
	Tasks      []Task
	Navigation string
	Search     string
	Message    string
}
