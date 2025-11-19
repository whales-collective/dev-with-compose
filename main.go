package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Human struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func textHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	human := Human{
		Name: "John Doe",
		Age:  30,
		City: "Paris",
	}

	fmt.Fprintf(w, "Name: %s\nAge: %d\nCity: %s", human.Name, human.Age, human.City)
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	human := Human{
		Name: "John Doe",
		Age:  30,
		City: "Paris",
	}

	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <title>Human Info</title>
</head>
<body>
    <h1>Human Information</h1>
    <p><strong>Name:</strong> %s</p>
    <p><strong>Age:</strong> %d</p>
    <p><strong>City:</strong> %s</p>
</body>
</html>
`, human.Name, human.Age, human.City)

	fmt.Fprint(w, html)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	human := Human{
		Name: "John Doe",
		Age:  30,
		City: "Paris",
	}

	json.NewEncoder(w).Encode(human)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>API Routes</title>
</head>
<body>
    <h1>Available Routes</h1>
    <ul>
        <li><a href="/text">/text</a> - Text response</li>
        <li><a href="/html">/html</a> - HTML response</li>
        <li><a href="/json">/json</a> - JSON response</li>
    </ul>
</body>
</html>
`
	fmt.Fprint(w, html)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	response := map[string]any{
		"status": "healthy",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/text", textHandler)
	http.HandleFunc("/html", htmlHandler)
	http.HandleFunc("/json", jsonHandler)
    http.HandleFunc("/health", healthCheckHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
