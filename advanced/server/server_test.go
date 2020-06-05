package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mactsouk/handlers"
)

func TestTimeHanlderV1(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/time", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.TimeHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestTimeHanlderV2(t *testing.T) {
	req, err := http.NewRequest("GET", "/v2/time", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.TimeHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
