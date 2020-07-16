package main

import (
	"fmt"
	"net/http"
)

func index(request http.ResponseWriter, response *http.Request) {
	fmt.Fprintf(request, "Hello World")
}

func about(request http.ResponseWriter, response *http.Request) {
	fmt.Fprintf(request, "About")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server started...")
	http.ListenAndServe(":3000", nil)
}
