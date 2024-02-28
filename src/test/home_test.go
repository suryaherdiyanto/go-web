package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	Setup()
	ts := httptest.NewServer(singleton.Session.LoadAndSave(routes))

	defer ts.Close()

	r, err := http.Get(ts.URL)

	if err != nil {
		t.Error(err)
	}

	if r.StatusCode != 200 {
		t.Errorf("Expected status code %d and got %d", 200, r.StatusCode)
	}
}
