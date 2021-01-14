package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1 style='color: green;'>Hello World</h1>")
}

func HandleProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<ul><li>laptop</li><li>Iphone</li><li>Xiaomi</li></ul>")
}
