# HTML TEMPLATES

Go’s html/template package provides a rich templating language for HTML templates. It is mostly used in web applications to display data in a structured way in a client’s browser. 

## FIRST TEMPLATE

to access variables we preappend the `.` i.e

```go
data := TodoPageData {
   PageTitle: "My Todo List ",
   Todos ; []Todo{
    {Title: "Watch Chelsea match", Done: true},
    {Title: "Get depressed about Chelsea", Done: true},
    {Title: "Watch another Chelsea match again", Done: false}
   },
}
```

then on the template:

```html
<h1>{{ .PageTitle}}</h1>
<ul>
    {{range .Todos}}
        {{ if .Done}}
            <li class = "Done">{{ .Title}}</li>
        {{ else }}
            <li>{{ .Title }}</li>
        {{ end }}
    {{ end }}
</ul>
```

## templating language

1. {{ /* template comments */ }} - this is used to define comments
2. {{ . }} - render the root element
3. {{ .Title }} - render the title field in a nested element
4. {{if .Done}} {{else}} {{end}}	Defines an if-Statement
5. {{range .Todos}} {{.}} {{end}}	Loops over all “Todos” and renders each using {{.}}
6. {{block "content" .}} {{end}}	Defines a block with the name “content”

## parsing templates from files

