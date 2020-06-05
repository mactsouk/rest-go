package main

import (
	"bytes"
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

func TestNotFoundV1(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/doesNotExist", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.DefaultHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestNotFoundV2(t *testing.T) {
	req, err := http.NewRequest("GET", "/v2/doesNotExist", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.DefaultHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestMethodNotAllowedV1(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/v1/time", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.MethodNotAllowedHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestMethodNotAllowedV2(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/v2/time", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.MethodNotAllowedHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetallV1(t *testing.T) {
	UserPass := []byte(`{"user": "admin", "password": "1"}`)
	req, err := http.NewRequest("GET", "/v1/getall", bytes.NewBuffer(UserPass))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetAllHandlerUpdated)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}

	expected := `[{"id":1,"user":"admin","password":"1","lastlogin":0,"admin":1,"active":0}]`
	if rr.Body.String() != expected+"\n" {
		t.Errorf("handler returned unexpected body: got %v but wanted %v", rr.Body.String(), expected)
	}
}

func TestGetallV2(t *testing.T) {
	UserPass := []byte(`{"username": "admin", "password": "1", "U": {}}`)
	req, err := http.NewRequest("GET", "/v2/getall", bytes.NewBuffer(UserPass))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetAllHandlerV2)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}

	expected := `[{"id":1,"user":"admin","password":"1","lastlogin":0,"admin":1,"active":0}]`
	if rr.Body.String() != expected+"\n" {
		t.Errorf("handler returned unexpected body: got %v but wanted %v",
			rr.Body.String(), expected)
	}
}
