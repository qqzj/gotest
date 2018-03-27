package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/welcome.html", welcome)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
func welcome(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("<h1>Welcome to Go World</h1>"))
}
func index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("<h1>Index</h1>"))
}
