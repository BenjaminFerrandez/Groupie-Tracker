package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//take the data from all API structs.
type Data struct {
	A Artist
	R Relation
	L Location
	D Date
}

//stores data from artist API struct.
type Artist struct {
	Id 			 int 				`json:"id"`
	Image        string   			`json:"image"`
	Name         string    			`json:"name"`
	Members      []string			`json:"members"`
	CreationDate int				`json:"creationDate"`
	FirstAlbum   string				`json:"firstAlbum"`
	Locations 	 Location			`json:"locations"`
	ConcertDates Date				`json:"concertDates"`
	Relations    Relation			`json:"relations"`
	Dates        [][]string
}

//stores data from location API struct.
type Location struct {
	Id 			 int				`json:"id"`
	Locations    []string			`json:"locations"`
}

//stores data from date API struct.
type Date struct {
	Id 			 int				`json:"id"`
	Dates 		 []string			`json:"dates"`
}

//stores data from relation API struct.
type Relation struct {
	Id			   int						`json:"id"`
	DatesLocations map[string][]string		`json:"datesLocations"`
}

type Text struct {
	ErrorNum int
	ErrorMes string
}

type Structure struct {
	artists		  []Artist
	artistsTemp   []Artist
	numberOfGroup int
	action		  string
}

type WebStruct struct {
	Artists []Artist
}

var (
	artistInfo   []Artist
	locationMap  map[string]json.RawMessage
	locationInfo []Location
	datesMap     map[string]json.RawMessage
	datesInfo    []Date
	relationMap  map[string]json.RawMessage
	relationInfo []Relation
)

//gets and stores data from Artist API
func ArtistData() []Artist {
	artist, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal()
	}
	artistData, err := ioutil.ReadAll(artist.Body)
	if err != nil {
		log.Fatal()
	}
	json.Unmarshal(artistData, &artistInfo)
	return artistInfo
}

//gets and stores data from Location API
func LocationData() []Location {
	var bytes []byte
	location, err2 := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err2 != nil {
		log.Fatal()
	}
	locationData, err3 := ioutil.ReadAll(location.Body)
	if err3 != nil {
		log.Fatal()
	}
	err := json.Unmarshal(locationData, &locationMap)
	if err != nil {
		fmt.Println("error :", err)
	}
	for _, m := range locationMap {
		for _, v := range m {
			bytes = append(bytes, v)
		}
	}
	err = json.Unmarshal(bytes, &locationInfo)
	if err != nil {
		fmt.Println("error :", err)
	}
	return locationInfo
}

//gets and stores data from Dates API
func DatesData() []Date {
	var bytes []byte
	dates, err2 := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err2 != nil {
		log.Fatal()
	}
	datesData, err3 := ioutil.ReadAll(dates.Body)
	if err3 != nil {
		log.Fatal()
	}
	err := json.Unmarshal(datesData, &datesMap)
	if err != nil {
		fmt.Println("error :", err)
	}
	for _, m := range datesMap {
		for _, v := range m {
			bytes = append(bytes, v)
		}
	}
	err = json.Unmarshal(bytes, &datesInfo)
	if err != nil {
		fmt.Println("error :", err)
	}
	return datesInfo
}

//gets and stores data from Relation API
func RelationData() []Relation {
	var bytes []byte
	relation, err2 := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err2 != nil {
		log.Fatal()
	}
	relationData, err3 := ioutil.ReadAll(relation.Body)
	if err3 != nil {
		log.Fatal()
	}
	err := json.Unmarshal(relationData, &relationMap)
	if err != nil {
		fmt.Println("error :", err)
	}

	for _, m := range relationMap {
		for _, v := range m {
			bytes = append(bytes, v)
		}
	}

	err = json.Unmarshal(bytes, &relationInfo)
	if err != nil {
		fmt.Println("error :", err)
	}
	return relationInfo
}

//take the data from all API slices into one data struct.
func collectData() []Data {
	ArtistData()
	RelationData()
	LocationData()
	DatesData()
	dataData := make([]Data, len(artistInfo))
	for i := 0; i < len(artistInfo); i++ {
		dataData[i].A = artistInfo[i]
		dataData[i].R = relationInfo[i]
		dataData[i].L = locationInfo[i]
		dataData[i].D = datesInfo[i]
	}
	return dataData
}
