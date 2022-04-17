package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	strings "strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // write data to response
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		//从项目根路径下读取文件
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		//post method
		r.ParseForm()
		// logic part of log in
		// r.Form is type of map[String][]String
		//fmt.Println("username:", r.Form["username"])
		//fmt.Println("password:", r.Form["password"])
		//cross site
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) // print at server side
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))

		//t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		//err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
	}
}

func main() {
	http.HandleFunc("/", sayhelloName) // setting router rule
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
