# HTTP SERVER

there is some bear minimum for an http server.

1. it should be abble to process dynamic requests
2. it should be able to serve static assets.
3. it should be able to accept connections

## process dynamic requests

process incoming requests from users who browse the website, and allow them to process the information theu need

the `net/http` contains all the functionalities required to handle this.

we can always register handlers using the `http.HandleFunc()` which takes two parameters:

1. the path requested for
2. the function handler to be called

## serve static sites

serve javascript files, css and images to browsers to create dynamic experiece to the user

to achieve this we are going to use `http.FileServer`, and point it to url path 

`fs :=http.FileServer(http.Dir("static/"))` - this is used to inform our server where to get our static files

`http.Handle("/static/", http.StripPrefix("/static/", fs))` - this is used to point a url path to it, the stripprefix is used to help us to use relative paths instead of absolute paths

## accepting connections

the http server must be able to listen to a specific connection/ port to accept connections from the internet

for it to accept connections from the  internet is pretty simple: `http.ListenAndServer()`

our http server lives inside the main(), and so we are going to place it inside the main(), I mean the three functionalities are going to be inside the main() function.