package main

import (
	"fmt"
	"net/http"
)

func web() {
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("template"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.Handle("/font/", http.StripPrefix("/font/", http.FileServer(http.Dir("font"))))
	http.HandleFunc("/", Home)

	serverAddr := "localhost" + PORT
	fmt.Println("Server started on:", serverAddr)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
