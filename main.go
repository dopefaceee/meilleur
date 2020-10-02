package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone visit our web")
	fmt.Fprint(w, "<h1>Hello dopefaceee, trying bro!</h1>")
}
