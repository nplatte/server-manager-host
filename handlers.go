package main

import (
	"net/http"
	"text/template"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	template, _ := template.ParseFiles("templates/home.html")
	template.Execute(w, nil)
}
