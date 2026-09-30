package main

import (
	"encoding/json"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type LocationResponse struct {
	Index []Location `json:"index"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DatesResponse struct {
	Index []Dates `json:"index"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationResponse struct {
	Index []Relation `json:"index"`
}

type ArtistDetails struct {
	Artist   Artist
	Relation Relation
}

func fetchArtists() ([]Artist, error) {

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var artists []Artist

	err = json.NewDecoder(resp.Body).Decode(&artists)

	return artists, err
}

func fetchLocations() ([]Location, error) {

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response LocationResponse

	err = json.NewDecoder(resp.Body).Decode(&response)

	return response.Index, err
}

func fetchDates() ([]Dates, error) {

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response DatesResponse

	err = json.NewDecoder(resp.Body).Decode(&response)

	return response.Index, err
}

func fetchRelations() ([]Relation, error) {

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response RelationResponse

	err = json.NewDecoder(resp.Body).Decode(&response)

	return response.Index, err
}
