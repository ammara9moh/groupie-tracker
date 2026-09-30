package main

import "testing"

func TestFetchArtists(t *testing.T) {
	artists, err := fetchArtists()

	if err != nil {
		t.Fatalf("fetchArtists() returned an error: %v", err)
	}

	if len(artists) == 0 {
		t.Fatal("fetchArtists() returned no artists")
	}

	if artists[0].Name == "" {
		t.Fatal("first artist has no name")
	}
}

func TestFetchLocations(t *testing.T) {
	locations, err := fetchLocations()

	if err != nil {
		t.Fatalf("fetchLocations() returned an error: %v", err)
	}

	if len(locations) == 0 {
		t.Fatal("fetchLocations() returned no locations")
	}

	if locations[0].ID == 0 {
		t.Fatal("first location has an invalid ID")
	}
}

func TestFetchDates(t *testing.T) {
	dates, err := fetchDates()

	if err != nil {
		t.Fatalf("fetchDates() returned an error: %v", err)
	}

	if len(dates) == 0 {
		t.Fatal("fetchDates() returned no dates")
	}

	if dates[0].ID == 0 {
		t.Fatal("first date has an invalid ID")
	}
}

func TestFetchRelations(t *testing.T) {
	relations, err := fetchRelations()

	if err != nil {
		t.Fatalf("fetchRelations() returned an error: %v", err)
	}

	if len(relations) == 0 {
		t.Fatal("fetchRelations() returned no relations")
	}

	if relations[0].ID == 0 {
		t.Fatal("first relation has an invalid ID")
	}
}
