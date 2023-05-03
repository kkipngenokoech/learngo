# HELLO MULTIVERSE SERVER

Go has a web server already installed to it - net/http

it has all the http protocol functionalities.

## registering a request handler

this will be used to receive all incoming http connections from the browser, http clients or api requests.

a handler looks like this: `func (w http.Responsewriter, r *http.Request)`

this function receives two parameters:
    1. http.responseWriter -  this is where you write your text/html response to
    2. http.requests - it contains all the information about this http request including the urls

a handler can also be redifined this way: `http.HandleFunc("/", handler)` where handle is a function defined elsewhere.

the syntax for this handler is as follows:

```go
func handler(w http.ResponseWriter, r *http.request){
    // syntax goes here
}
```

## listen and server

we have to listen to the port for connection requests and we are going to do this using http.ListenAndServe(":port number", nil), this http server will listen to the port and pass the request to the handler.

## to start a go server

to start the local host server for go you are going to use `go run <name of your go file>`
