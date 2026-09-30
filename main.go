package main

import (
	"fmt"
	"net/http"
)

func main() {

	var err error

	artists, err = fetchArtists()
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		return
	}

	locations, err = fetchLocations()
	if err != nil {
		fmt.Println("Error fetching locations:", err)
		return
	}

	dates, err = fetchDates()
	if err != nil {
		fmt.Println("Error fetching dates:", err)
		return
	}

	relations, err = fetchRelations()
	if err != nil {
		fmt.Println("Error fetching relations:", err)
		return
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/artist", artistHandler)
	http.HandleFunc("/search", searchHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
