package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("WEATHER_API_KEY is not set")
	}

	response, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=Istanbul&aqi=no", apiKey))
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}

	defer response.Body.Close()
	if response.StatusCode > 299 {
		body, _ := httputil.DumpResponse(response, true)
		log.Fatalf("Response failed with status code: %d and body: %s", response.StatusCode, body)
	}

	var apiResponse map[string]any
	fmt.Printf("%T\n", apiResponse)
	fmt.Println(apiResponse)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	fmt.Println(apiResponse)

	location, ok := apiResponse["location"].(map[string]any)
	if !ok {
		fmt.Println("\nCannot assert into location")
	}
	current, ok := apiResponse["current"].(map[string]any)
	if !ok {
		fmt.Println("Cannot assert into current")
	}
	condition, ok := current["condition"].(map[string]any)
	if !ok {
		fmt.Println("Cannot assert into condition")
	}
	fmt.Printf("\n\n\n%T\n\n\n", condition)

	fmt.Printf("Location: %s, %s\nTemperature: %.1f°C\nCondition: %s\n\n", location["name"], location["country"], current["temp_c"].(float64), condition["text"])

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Location: %s, %s\nTemperature: %.1f°C\nCondition: %s\nLast Updated: %s\nUrl Path: %q", location["name"], location["country"], current["temp_c"].(float64), condition["text"], current["last_updated"], html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
