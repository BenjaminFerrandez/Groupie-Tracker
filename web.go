package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

//handles the artist Page when artist image is clicked by receiving "ArtistName" value
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

//search bar (doesn't work :( )
func (g *Structure) searchGroup(str string) {
	g.artistsTemp = []Artist{}
	for i := 0; i < len(g.artists); i++ {
		if strings.Contains(strings.ToLower(g.artists[i].Name), strings.ToLower(str)) {
			g.artistsTemp = append(g.artistsTemp, g.artists[i])
		}
		for j := 0; j < len(g.artists[i].Members); j++ {
			if strings.Contains(strings.ToLower(g.artists[i].Members[j]), strings.ToLower(str)) {
				g.artistsTemp = append(g.artistsTemp, g.artists[i])
			}
		}
	}
}

func (g *Structure) index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("pages/index.html"))
	r.ParseForm()
	action := r.Form.Get("action")
	g.action = action

	search := r.Form.Get("search")
	if len(search) > 0 {
		g.searchGroup(search)
	}

	web := WebStruct{Artists: g.artistsTemp}
	err := tmpl.Execute(w, web)
	if err != nil {
		return
	}
}