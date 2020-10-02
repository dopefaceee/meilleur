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
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Println("Someone visit our web")
		fmt.Fprint(w, "<h1>Hello dopefaceee, trying bro!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send me an email to<a href=\"mailto:ade.suluh@gmail.cp,\">ade.suluh@gmail.com</a>")

	}

}
