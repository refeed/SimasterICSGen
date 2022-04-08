package main

import (
	"fmt"
	"net/http"
)

const (
	HTTP_PORT = 8080
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Printf("Listening on port %v\n", HTTP_PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", HTTP_PORT), nil)
}
