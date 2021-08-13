package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/sid", SID)
	http.ListenAndServe(":8080", nil)
}
