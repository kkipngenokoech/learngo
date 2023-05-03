package main

import (
	"fmt"
	"net/http"
)

func main() {
	// handling dynamic user requests
	http.HandleFunc("/index", indexHandler)
	println("the server is running on port 8080")

	// handling static files
	fs := http.FileServer(http.Dir("/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs)

	// listening to the port server
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "welcome to my go server\n")
}