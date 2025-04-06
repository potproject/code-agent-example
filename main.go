package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// OpenWeatherMap APIのレスポンスから必要なデータだけを格納する構造体
type WeatherData struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

// ロサンゼルスの天気を取得して表示するプログラム
func main() {
	// ロサンゼルスの都市ID
	cityID := "5368361"
	// APIキー（無料で取得可能）
	apiKey := "YOUR_API_KEY"
	// OpenWeatherMap APIのURL
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?id=%s&appid=%s&units=metric", cityID, apiKey)

	// APIリクエスト
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("天気情報を取得できませんでした:", err)
		return
	}
	defer resp.Body.Close()

	// レスポンスの読み込み
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("レスポンスの読み込みに失敗しました:", err)
		return
	}

	// JSONデータをパース
	var data WeatherData
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("JSONデータの解析に失敗しました:", err)
		return
	}

	// 天気情報の表示
	if len(data.Weather) > 0 {
		fmt.Println("ロサンゼルスの今日の天気:")
		fmt.Printf("天気: %s\n", data.Weather[0].Description)
		fmt.Printf("気温: %.1f°C\n", data.Main.Temp)
		fmt.Printf("体感温度: %.1f°C\n", data.Main.FeelsLike)
		fmt.Printf("湿度: %d%%\n", data.Main.Humidity)
		fmt.Printf("風速: %.1fm/s\n", data.Wind.Speed)
	} else {
		fmt.Println("天気情報がありません")
	}
}
