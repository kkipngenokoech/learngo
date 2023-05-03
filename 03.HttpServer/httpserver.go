package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/index", indexHandler)
	println("the server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "welcome to my go server\n")
}