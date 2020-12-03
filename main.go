package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/about",aboutHandler)
	router.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprint(w, "访问文章列表1")
		case "POST":
			fmt.Fprint(w, "创建新的文章")
		}
	})
	http.ListenAndServe(":2000", router)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>欢迎来到我的博客</h1>   <a href='/about'>关于我 </a>  <a href='/articles'>文章 </a>")
	} else {
		fmt.Fprint(w,"<h1>页面未找到 ):  </h1>")
	}
}

func aboutHandler(w http.ResponseWriter ,r * http.Request){
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprint(w,"这个是关于博客的介绍2")
}
