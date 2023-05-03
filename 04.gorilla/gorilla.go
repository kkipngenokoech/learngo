package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/index", indexHandler)
	r.HandleFunc("/books/{title}/page/{page}", bookPageHandler)
	println("the server is running on port 8080")
	http.ListenAndServe(":8000", nil)
}


func indexHandler(w http.ResponseWriter, r *http.Request){
	println("we are her")
	fmt.Fprintf(w,"Hey there gorilla package is working")
}

func bookPageHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You are requesting a book: %s\n, page number %d\n", title, page)
}