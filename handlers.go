package main

import (
	"bytes"
	"embed"
	"io/fs"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

var artists []Artist
var locations []Location
var dates []Dates
var relations []Relation

//go:embed templates/*.html static/*
var siteFiles embed.FS

// Templates and static assets are embedded so they are included in Vercel's binary.
var templates = template.Must(template.ParseFS(siteFiles, "templates/*.html"))

func staticFiles() http.FileSystem {
	files, err := fs.Sub(siteFiles, "static")
	if err != nil {
		panic(err)
	}
	return http.FS(files)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := strings.TrimSpace(r.URL.Query().Get("name"))
	if name == "" {
		renderError(w, "Search name is required", http.StatusBadRequest)
		return
	}

	var results []Artist
	for _, artist := range artists {
		if strings.EqualFold(artist.Name, name) {
			results = append(results, artist)
		}
	}

	render(w, "search.html", results)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		renderError(w, "Page not found", http.StatusNotFound)
		return
	}

	render(w, "index.html", artists)
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
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
		renderError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	details := ArtistDetails{
		Artist:   selectedArtist,
		Relation: selectedRelation,
	}

	render(w, "artist.html", details)
}

// render builds the page in a buffer so a template failure can still return a clean 500.
func render(w http.ResponseWriter, name string, data interface{}) {
	var buf bytes.Buffer
	if err := templates.ExecuteTemplate(&buf, name, data); err != nil {
		renderError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}

// renderError sends the status code first, then the error page.
func renderError(w http.ResponseWriter, message string, status int) {
	var buf bytes.Buffer
	if err := templates.ExecuteTemplate(&buf, "error.html", message); err != nil {
		http.Error(w, message, status)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	buf.WriteTo(w)
}
