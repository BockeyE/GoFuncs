package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("server ready")
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}
