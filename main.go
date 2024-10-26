package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func getData(apiKey string, latitude string, longitude string) string {
	// Get data from API
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s", latitude, longitude, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error getting data from API")
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body")
	}

	// Convert body to string
	return string(body)
}

func parseData(data string) {
	// Parse data from API

	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Println("Error parsing data")
	}
	fmt.Println(m)
	n := m["main"].(map[string]interface{})

	fmt.Println(n["humidity"])

}

func scedule() {
	// Scedule the function to run every 5 minutes

}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	latitude := os.Getenv("LATITUDE")
	longitude := os.Getenv("LONGITUDE")

	data := getData(apiKey, latitude, longitude)
	parseData(data)

}
