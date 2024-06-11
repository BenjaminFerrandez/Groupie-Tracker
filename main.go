package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	HandleRequests()
}

//home page handler which executes the index.html file.
func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Println("Endpoint Hit: returnAllArtists")
	data := ArtistData()
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}