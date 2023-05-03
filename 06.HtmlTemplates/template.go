package main

import (

	"html/template"

	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct{
	Title string
	Done bool
}

type TodoPageData struct {
	PageTitle string
	Todos []Todo
}

func main() {
	router := mux.NewRouter()
	template := template.Must(template.ParseFiles("layout.html"))

	router.HandleFunc("/todo", todoHandler)
	http.ListenAndServe(":8080", router)
}

func todoHandler(w http.ResponseWriter, r *http.Request){
	data := TodoPageData{
		PageTitle: "Kip's Todo List",
		Todos: []Todo{
			{Title: "Watch Chelsea match", Done: true},
    		{Title: "Get depressed about Chelsea", Done: true},
    		{Title: "Watch another Chelsea match again", Done: false},
		},
	}
	template.Execute(w, data)
}