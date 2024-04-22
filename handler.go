package main

import (
	"html/template"
	"net/http"
)


func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func renderTemplate(w http.ResponseWriter, html string) {
	t, err := template.ParseFiles("./template/" + html + ".page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
