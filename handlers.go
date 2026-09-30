package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var artists []Artist
var locations []Location
var dates []Dates
var relations []Relation

func searchHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	name := strings.TrimSpace(r.URL.Query().Get("name")) //used to read the search item, it reads the value of the name query parameter
	//TrimSpace is used to remove spaces

	if name == "" {
		http.Error(w, "Search name is required", http.StatusBadRequest)
		return
	}

	var results []Artist

	for _, artist := range artists {
		if strings.EqualFold(artist.Name, name) { //EqualFold compares strings without caring about uppercase/lowercase.
			results = append(results, artist)
		}
	}

	tmpl, err := template.ParseFiles("templates/search.html")

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, results)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	if r.URL.Path != "/" {
		renderError(w, "Page not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artists)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func artistHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		renderError(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	var selectedArtist Artist
	foundArtist := false

	for _, artist := range artists {
		if artist.ID == id {
			selectedArtist = artist
			foundArtist = true
			break
		}
	}

	if !foundArtist {
		renderError(w, "Artist not found", http.StatusNotFound)
		return
	}

	var selectedRelation Relation
	foundRelation := false

	for _, relation := range relations {
		if relation.ID == selectedArtist.ID {
			selectedRelation = relation
			foundRelation = true
			break
		}
	}

	if !foundRelation {
		renderError(w, "Artist relation not found", http.StatusNotFound)
		return
	}

	details := ArtistDetails{
		Artist:   selectedArtist,
		Relation: selectedRelation,
	}

	tmpl, err := template.ParseFiles("templates/artist.html")

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, details)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func renderError(w http.ResponseWriter, message string, status int) {

	tmpl, err := template.ParseFiles("templates/error.html")

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, message)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
}
