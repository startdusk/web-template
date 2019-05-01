package main

import (
	"html/template"
	"log"
	"net/http"
)

// Profile 列表属性
type Profile struct {
	Name    string
	Work    string
	Amateur string
}

// items 定义列表内容
var items = []Profile{{
	Name:    "Benjamin",
	Work:    "PHP",
	Amateur: "Go",
}, {
	Name:    "Shelby",
	Work:    "Java",
	Amateur: "PHP",
}, {
	Name:    "John",
	Work:    "C++",
	Amateur: "Ruby",
}, {
	Name:    "Marry",
	Work:    "C#",
	Amateur: "JavaScript",
}}

// isOdd 单数为false，双数为true
func isOdd(num int) bool {
	return (num % 2) == 0
}

// index 首页句柄函数
func index(w http.ResponseWriter, r *http.Request) {
	var (
		tmpl *template.Template
		err  error
	)

	tmpl, err = template.New("index.html").
		Funcs(template.FuncMap{"isOdd": isOdd}).
		ParseFiles("./views/index.html")
	if err != nil {
		log.Println("parse template file error:", err.Error())
		return
	}

	err = tmpl.Execute(w, items)
	if err != nil {
		log.Println("template execute error:", err.Error())
		return
	}
}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
