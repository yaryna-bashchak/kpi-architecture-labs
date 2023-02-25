package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/time", Handler)
	http.ListenAndServe(":8795", nil)
}