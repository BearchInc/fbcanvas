package main

import (
	"net/http"
	"html/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("public/index.html")
	data := struct{User string}{"Diego"}
	t.Execute(w, data)
}

func init() {
	http.HandleFunc("/app/", handler)
}
