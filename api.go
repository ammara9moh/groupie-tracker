package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client = http.Client{Timeout: 10 * time.Second}

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
	var artists []Artist
	err := fetchJSON("https://groupietrackers.herokuapp.com/api/artists", &artists)
	return artists, err
}

func fetchLocations() ([]Location, error) {
	var response LocationResponse
	err := fetchJSON("https://groupietrackers.herokuapp.com/api/locations", &response)
	return response.Index, err
}

func fetchDates() ([]Dates, error) {
	var response DatesResponse
	err := fetchJSON("https://groupietrackers.herokuapp.com/api/dates", &response)
	return response.Index, err
}

func fetchRelations() ([]Relation, error) {
	var response RelationResponse
	err := fetchJSON("https://groupietrackers.herokuapp.com/api/relation", &response)
	return response.Index, err
}

func fetchJSON(url string, target interface{}) error {
	resp, err := client.Get(url) //10 seconds is for the whole request: connecting, sending, waiting for the answer, and reading the body.
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s returned status %d", url, resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
