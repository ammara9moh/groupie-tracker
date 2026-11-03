package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(staticFiles())))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	address := ":" + port
	fmt.Printf("Server running on http://localhost%s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
