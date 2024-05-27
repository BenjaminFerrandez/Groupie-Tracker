package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

//handles the artist Page when artist image is clicked by receiving "ArtistName" value
// and comparing it to the names in Data.Artist.Name field.
// Tells server what enpoints users hit.
func artistPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artistInfo" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Println("Endpoint Hit: Artist's Page")
	value := r.FormValue("ArtistName")
	if value == "" {
		errorHandler(w, r, http.StatusBadRequest)
		return
	}
	a := collectData()
	var b Data
	for i, ele := range collectData() {
		if value == ele.A.Name {
			b = a[i]
		}
	}
	t, err := template.ParseFiles("template/artistPage.html")
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
	t.Execute(w, b)
}

// Tells server what enpoints users hit.
// displays location data as a JSON raw message on webpage.
func returnAllLocations(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllLocations")
	json.NewEncoder(w).Encode(LocationData())
}

// Tells server what enpoints users hit.
// displays dates data as a JSON raw message on webpage.
func returnAllDates(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllDates")
	json.NewEncoder(w).Encode(DatesData())
}

// Tells server what enpoints users hit.
// displays relation data as a JSON raw message on webpage.
func returnAllRelation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllRelation")
	json.NewEncoder(w).Encode(RelationData())
}