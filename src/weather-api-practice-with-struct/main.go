package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"
)

type WeatherResponse struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		LastUpdated string  `json:"last_updated"`
		TempC       float64 `json:"temp_c"`
		Condition   struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

type Weather struct {
	City        string  `json:"city"`
	Country     string  `json:"country"`
	Temperature float64 `json:"temperature"`
	Condition   string  `json:"condition"`
	LastUpdated string  `json:"last_updated"`
}

func main() {
	// Get the API key from the environment variable
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("WEATHER_API_KEY is not set")
	}

	// Make the API request
	response, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=Istanbul&aqi=no", apiKey))
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		body, _ := io.ReadAll(response.Body)
		log.Fatalf("Response failed with status code: %d and body: %s", response.StatusCode, body)
	}

	var apiResponse WeatherResponse
	fmt.Printf("%T\n", apiResponse)
	fmt.Println(apiResponse)
	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	fmt.Println(apiResponse)

	// Map the data to the Weather struct
	weather := Weather{
		City:        apiResponse.Location.Name,
		Country:     apiResponse.Location.Country,
		Temperature: apiResponse.Current.TempC,
		Condition:   apiResponse.Current.Condition.Text,
		LastUpdated: apiResponse.Current.LastUpdated,
	}

	// Display the weather info
	fmt.Printf("City: %s\nTemperature: %.1f°C\nCondition: %s\n Country:%s\n", weather.City, weather.Temperature, weather.Condition, weather.Country)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "City: %s\nTemperature: %.1f°C\nCondition: %s\nCountry: %s\nLast Updated: %s\n%q\n", weather.City, weather.Temperature, weather.Condition, weather.Country, weather.LastUpdated, html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
