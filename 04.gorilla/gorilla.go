package main

import (
	"fmt"
	"net/http"
	
	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/books/{title}/page/{page}", bookPageHandler)
	println("the server is running on port 8080")
	http.ListenAndServe(":8080", router)
}


func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hey there gorilla package is working")
}

func bookPageHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	title := vars["title"]
	page :=  vars["page"]

	fmt.Fprintf(w, "You are requesting a book: %s, page number %s\n", title, page)
}