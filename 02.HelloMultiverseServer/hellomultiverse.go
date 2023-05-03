package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Println("server running on port 8080")
		fmt.Fprintf(w, "Hello and welcome to multiverse, you request for: %s\n", r.URL.Path)
	})
	http.HandleFunc("/show", Show)
	http.HandleFunc("/captain", Captain)
	http.ListenAndServe(":8080",nil)
}

func Show(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hi there! let me show you how amazing you are because you requested for: %s\n", r.URL.Path)
}

func Captain(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hi there captain, welcome to the deck")
}