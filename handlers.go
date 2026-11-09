package main

import (
	"encoding/json"
	"net/http"
	"text/template"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	template, _ := template.ParseFiles("templates/home.html")
	template.Execute(w, nil)
}

func HandleViewDevicesAll(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode([]string{})
}
