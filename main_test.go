package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*
go test -v
go test -cover
*/

func TestTextHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/text", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(textHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Name: John Doe\nAge: 30\nCity: Paris"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	contentType := rr.Header().Get("Content-Type")
	if contentType != "text/plain" {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "text/plain")
	}
}

func TestHtmlHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/html", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(htmlHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "John Doe") {
		t.Error("handler response should contain 'John Doe'")
	}
	if !strings.Contains(body, "30") {
		t.Error("handler response should contain '30'")
	}
	if !strings.Contains(body, "Paris") {
		t.Error("handler response should contain 'Paris'")
	}
	if !strings.Contains(body, "<html>") {
		t.Error("handler response should contain HTML tags")
	}

	contentType := rr.Header().Get("Content-Type")
	if contentType != "text/html" {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "text/html")
	}
}

func TestJsonHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/json", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(jsonHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var human Human
	err = json.Unmarshal(rr.Body.Bytes(), &human)
	if err != nil {
		t.Errorf("handler returned invalid JSON: %v", err)
	}

	if human.Name != "John Doe" {
		t.Errorf("unexpected Name: got %v want %v", human.Name, "John Doe")
	}
	if human.Age != 30 {
		t.Errorf("unexpected Age: got %v want %v", human.Age, 30)
	}
	if human.City != "Paris" {
		t.Errorf("unexpected City: got %v want %v", human.City, "Paris")
	}

	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "application/json")
	}
}

func TestRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rootHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "/text") {
		t.Error("handler response should contain '/text' route")
	}
	if !strings.Contains(body, "/html") {
		t.Error("handler response should contain '/html' route")
	}
	if !strings.Contains(body, "/json") {
		t.Error("handler response should contain '/json' route")
	}
	if !strings.Contains(body, "<html>") {
		t.Error("handler response should contain HTML tags")
	}

	contentType := rr.Header().Get("Content-Type")
	if contentType != "text/html" {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "text/html")
	}
}

func TestHumanStruct(t *testing.T) {
	human := Human{
		Name: "Test User",
		Age:  25,
		City: "London",
	}

	if human.Name != "Test User" {
		t.Errorf("unexpected Name: got %v want %v", human.Name, "Test User")
	}
	if human.Age != 25 {
		t.Errorf("unexpected Age: got %v want %v", human.Age, 25)
	}
	if human.City != "London" {
		t.Errorf("unexpected City: got %v want %v", human.City, "London")
	}
}

func TestHumanJsonMarshalling(t *testing.T) {
	human := Human{
		Name: "Jane Smith",
		Age:  28,
		City: "Berlin",
	}

	jsonData, err := json.Marshal(human)
	if err != nil {
		t.Fatalf("failed to marshal Human: %v", err)
	}

	var unmarshalled Human
	err = json.Unmarshal(jsonData, &unmarshalled)
	if err != nil {
		t.Fatalf("failed to unmarshal Human: %v", err)
	}

	if unmarshalled.Name != human.Name {
		t.Errorf("Name mismatch after marshal/unmarshal: got %v want %v", unmarshalled.Name, human.Name)
	}
	if unmarshalled.Age != human.Age {
		t.Errorf("Age mismatch after marshal/unmarshal: got %v want %v", unmarshalled.Age, human.Age)
	}
	if unmarshalled.City != human.City {
		t.Errorf("City mismatch after marshal/unmarshal: got %v want %v", unmarshalled.City, human.City)
	}
}
