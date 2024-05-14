package main

import (
	"bytes"
	"github.com/Sal1mbek/letsgofurther/internal/data"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestUpdateMovieHandler(t *testing.T) {
	// Mock application and movie data
	id := 1
	expectedTitle := "Updated Title"
	expectedYear := 2024
	expectedRuntime := data.Runtime{Minutes: 120}
	expectedGenres := []string{"action", "adventure"}

	movie := &data.Movie{ID: id, Title: "Original Title", Year: 2023, Runtime: data.Runtime{Minutes: 90}, Genres: []string{"action"}}
	updatedMovie := &data.Movie{ID: id, Title: expectedTitle, Year: expectedYear, Runtime: expectedRuntime, Genres: expectedGenres}

	app := &application{
		models: &mockModels{
			GetFunc: func(id int) (*data.Movie, error) {
				return movie, nil
			},
			UpdateFunc: func(m *data.Movie) error {
				*m = *updatedMovie
				return nil
			},
		},
	}

	// Create a request with a valid ID and JSON body for update
	reqBody := []byte(`{"title": "` + expectedTitle + `", "year": ` + strconv.Itoa(expectedYear) + `, "runtime": "` + expectedRuntime.String() + `", "genres": ["` + strings.Join(expectedGenres, `", "`) + `"]}`)
	req, err := http.NewRequest(http.MethodPatch, "/v1/movies/"+strconv.Itoa(id), bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Set up recorder to capture response
	rr := httptest.NewRecorder()

	// Call the handler function
	updateMovieHandler(app, rr, req)

	// Check for expected status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// No need to check response body as the success is indicated by status code

	// Verify the movie was updated in the mock model
	if app.models.(*mockModels).Get(id).Title != expectedTitle {
		t.Errorf("expected movie title to be updated in mock model")
	}
}
