# Gorilla/mux

this package helps us to write named paths/urls.

we are going to use it to create routes with named parameters, GET/POST handlers with domain restrictions

to install this package you: `go get -u github.com/gorilla/mux`

you might need to create a go module for this particular project: `go mod init <name of the module>`

to use this package in your code all you need to do is to import it: `import "github.com/gorilla/mux"`

## creating a new router

to create a new router using the mux package: `index := mux.NewRouter()`

this will receive all the requests and pass them to the handlers you would have defined.

using mux to register handlers: `index.HandleFunc()`

## dyanamic routes

to create a new route one has to i.e `/books/The-Alchemist/page/10` in this url path we have to dynamic sites i.e The-Alchemist and 10.

to implement this:

`index.HandleFunc("/books/{title}/page/{page}", getPageHandler)`

then in the getPageHandler:

```go
func getPageHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.vars(r)
    vars["title"]
    vars["page"]
}
```

we would need to extract this url parameters and to do that we are going to use the `vars := mux.vars[r]`

it takes http.request as a parameter and returns a map of the segments.
