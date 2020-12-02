package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/2", handlerFunc2)
	http.ListenAndServe(":3000", nil)
}

func handlerFunc(W http.ResponseWriter, r *http.Request) {
	fmt.Fprint(W, "<h1>Hello,world</h1>")
	r.Header.Get("User-Agent")
}

func handlerFunc2(W http.ResponseWriter, r *http.Request) {
	fmt.Fprint(W, "<h1>Hello,world2</h1>")
}