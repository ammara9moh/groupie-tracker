package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	homeHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d",
			http.StatusOK, rec.Code)
	}

}

func TestHomeHandlerMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()

	homeHandler(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %d, got %d",
			http.StatusMethodNotAllowed, rec.Code)
	}

}

func TestHomeHandlerNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	rec := httptest.NewRecorder()

	homeHandler(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d",
			http.StatusNotFound, rec.Code)
	}

}

func TestArtistHandlerInvalidID(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/artist?id=abc", nil)
	rec := httptest.NewRecorder()

	artistHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d",
			http.StatusBadRequest, rec.Code)
	}

}

func TestArtistHandlerNotFound(t *testing.T) {
	oldArtists := artists
	defer func() { artists = oldArtists }()

	artists = []Artist{
		{ID: 1, Name: "Test Artist"},
	}

	req := httptest.NewRequest(http.MethodGet, "/artist?id=999", nil)
	rec := httptest.NewRecorder()

	artistHandler(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d",
			http.StatusNotFound, rec.Code)
	}

}

func TestArtistHandlerMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/artist?id=1", nil)
	rec := httptest.NewRecorder()

	artistHandler(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %d, got %d",
			http.StatusMethodNotAllowed, rec.Code)
	}

}

// Simulates a GET request to /search without a name parameter.
// Runs searchHandler.
// Checks that the handler returns 400 Bad Request, because the search name is required.
func TestSearchHandlerMissingName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/search", nil)
	rec := httptest.NewRecorder()

	searchHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", rec.Code)
	}

}

// Creates test data containing an artist named Queen.
// Sends a search request for Queen.
// Runs searchHandler.
// Checks that the response is 200 OK.
func TestSearchHandlerSuccess(t *testing.T) {
	oldArtists := artists
	defer func() { artists = oldArtists }()

	artists = []Artist{
		{ID: 1, Name: "Queen"},
	}

	req := httptest.NewRequest(http.MethodGet, "/search?name=Queen", nil)
	rec := httptest.NewRecorder()

	searchHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rec.Code)
	}

}

// Searches for queen using lowercase letters.
// Checks that the response is 200 OK.
// Checks that the returned HTML contains Queen.
func TestSearchHandlerCaseInsensitive(t *testing.T) {
	oldArtists := artists
	defer func() { artists = oldArtists }()

	artists = []Artist{
		{ID: 1, Name: "Queen"},
	}

	req := httptest.NewRequest(http.MethodGet, "/search?name=queen", nil)
	rec := httptest.NewRecorder()

	searchHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "Queen") {
		t.Error("Expected search results to contain Queen")
	}

}

// Sends a POST request to /search.
// Your handler only allows GET.
// Checks that the response is 405 Method Not Allowed.
func TestSearchHandlerMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/search?name=Queen", nil)
	rec := httptest.NewRecorder()

	searchHandler(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected 405, got %d", rec.Code)
	}

}

// Calls renderError() with a 404 status.
// Checks whether the response actually returns 404 Not Found.
func TestRenderError(t *testing.T) {
	rec := httptest.NewRecorder()

	renderError(rec, "Page not found", http.StatusNotFound)

	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", rec.Code)
	}

}

// It verifies that renderError() returns HTTP status 500 when you pass http.StatusInternalServerError
func TestRenderErrorInternalServerError(t *testing.T) {
	rec := httptest.NewRecorder()

	renderError(rec, "Internal Server Error", http.StatusInternalServerError)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected 500, got %d", rec.Code)
	}

}
