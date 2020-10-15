package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// http/template

// 模板引擎的使用

func sayHello(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./bingo/lesson_16/hello.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	tmpl.Execute(w, "Golang")
}

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

// 模板语法
func sayStruct(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./bingo/lesson_16/stu.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := UserInfo{
		Name:   "kk",
		Gender: "男",
		Age:    23,
	}
	tmpl.Execute(w, user)
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/stu", sayStruct)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
