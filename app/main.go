package main

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates/*
var templatesFS embed.FS

//go:embed static/*
var staticFS embed.FS

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFS(templatesFS, "templates/index.html"))
	data := struct{ Title string }{Title: "Welcome to the Multi-Tenancy March Series Powered by Loft Labs"}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.FileServer(http.FS(staticFS)))
	http.ListenAndServe(":8080", nil)
}
