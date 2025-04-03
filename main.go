package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type WeatherResponse struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Name string `json:"name"`
}

// 東京の天気を取得して表示するプログラム
func main() {
	fmt.Println("Current time:", time.Now())
	
	weather, err := getTokyoWeather()
	if err != nil {
		fmt.Println("Error getting weather:", err)
		return
	}
	
	fmt.Printf("天気情報 (東京): %s (%.1f°C, 湿度: %d%%)\n", 
		weather.Weather[0].Description,
		weather.Main.Temp,
		weather.Main.Humidity)
}

func getTokyoWeather() (*WeatherResponse, error) {
	url := "https://api.openweathermap.org/data/2.5/weather?q=Tokyo&appid=API_KEY&units=metric&lang=ja"
	
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, err
	}
	
	return &weather, nil
}
