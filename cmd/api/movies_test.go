package main

import (
	"net/http"
	"testing"
)

func TestShowMovieHandler(t *testing.T) {
	app := newTestApp()

	ts := newTestServer(app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/v1/movies/1")

	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}

	expResp := `{
	"error": "you must be authenticated to access this resource"}
	`

	if string(body) != expResp {
		t.Errorf("want body to equal %q,\n but got %q", expResp, string(body))
	}
}
