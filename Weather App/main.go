package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func main() {
	var city string
	fmt.Print("Enter city name: ")
	fmt.Scanln(&city)

	apiKey := "de47fcee8782e4d45e7cdf7e2ef95ace"

	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city, apiKey,
	)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching weather data")
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println("City not found or API error")
		return
	}

	var weather WeatherResponse
	json.NewDecoder(response.Body).Decode(&weather)

	fmt.Println("\n🌍 City:", weather.Name)
	fmt.Println("🌡️ Temperature:", weather.Main.Temp, "°C")
	fmt.Println("💧 Humidity:", weather.Main.Humidity, "%")
	fmt.Println("☁️ Condition:", weather.Weather[0].Description)
}
