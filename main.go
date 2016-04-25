package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!\n\n")
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
