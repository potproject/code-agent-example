package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// OpenWeatherMap APIのレスポンス構造体
type WeatherResponse struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

// ロサンゼルスの天気を取得して表示するプログラム
func main() {
	// ロサンゼルスの天気を取得
	weather, err := getLAWeather()
	if err != nil {
		fmt.Println("天気情報の取得に失敗しました:", err)
		return
	}

	// 結果を表示
	fmt.Printf("日時: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("場所: %s\n", weather.Name)
	if len(weather.Weather) > 0 {
		fmt.Printf("天気: %s\n", weather.Weather[0].Description)
	}
	fmt.Printf("気温: %.1f°C\n", weather.Main.Temp)
}

// ロサンゼルスの天気情報を取得する関数
func getLAWeather() (*WeatherResponse, error) {
	// OpenWeatherMap APIのエンドポイント (注: 実際の使用にはAPIキーが必要です)
	url := "https://api.openweathermap.org/data/2.5/weather?q=Los%20Angeles&units=metric&appid=YOUR_API_KEY"

	// HTTPリクエスト
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// レスポンスの読み取り
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// JSONのデコード
	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, err
	}

	return &weather, nil
}
