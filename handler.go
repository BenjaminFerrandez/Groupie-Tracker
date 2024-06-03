package main

import (
	"html/template"
	"net/http"
	"fmt"
)

func renderTemplate(w http.ResponseWriter, html string) {
	t, err := template.ParseFiles("./template/" + html + ".page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

//handles error messages
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		t, err := template.ParseFiles("errorPage.html")
		if err != nil {
			errorHandler(w, r, http.StatusInternalServerError)
			return
		}
		em := "HTTP status 404: Page Not Found"
		p := Text{ErrorNum: status, ErrorMes: em}
		t.Execute(w, p)
	}
	if status == http.StatusInternalServerError {
		t, err := template.ParseFiles("template/errorPage.html")
		if err != nil {
			fmt.Fprint(w, "HTTP status 500: Internal Server Error -missing errorPage.html file")
		}
		em := "HTTP status 500: Internal Server Error"
		p := Text{ErrorNum: status, ErrorMes: em}
		t.Execute(w, p)
	}
	if status == http.StatusBadRequest {
		t, err := template.ParseFiles("template/errorPage.html")
		if err != nil {
			fmt.Fprint(w, "HTTP status 500: Internal Server Error -missing errorPage.html file")
		}
		em := "HTTP status 400: Bad Request! Please select artist from the Home Page"
		p := Text{ErrorNum: status, ErrorMes: em}
		t.Execute(w, p)
	}
}

// collection of webpage handlers
func HandleRequests() {
	fmt.Println("Starting Server at Port 8080")
	fmt.Println("now open a broswer and enter 'localhost:8080' into the URL")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/artistInfo", artistPage)
	http.HandleFunc("/locations", returnAllLocations)
	http.HandleFunc("/dates", returnAllDates)
	http.HandleFunc("/relation", returnAllRelation)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.ListenAndServe(":8080", nil)
}