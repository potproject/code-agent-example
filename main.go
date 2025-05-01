package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// 天気情報を格納する構造体
type WeatherResponse struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Name string `json:"name"`
}

// 天気の状態を日本語に変換
func translateWeather(condition string) string {
	weatherMap := map[string]string{
		"Clear":        "晴れ",
		"Clouds":       "曇り",
		"Rain":         "雨",
		"Drizzle":      "小雨",
		"Thunderstorm": "雷雨",
		"Snow":         "雪",
		"Mist":         "霧",
		"Fog":          "霧",
		"Haze":         "もや",
		"Smoke":        "煙霧",
		"Dust":         "砂塵",
		"Sand":         "砂",
		"Ash":          "火山灰",
		"Squall":       "暴風",
		"Tornado":      "竜巻",
	}

	if japaneseWeather, ok := weatherMap[condition]; ok {
		return japaneseWeather
	}
	return condition
}

// 世界各国の時計と天気を表示するgo言語のプログラム
func main() {
	fmt.Println("=== 世界各国の時計と天気 ===")
	
	// 環境変数からAPIキーを取得
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		fmt.Println("警告: OPENWEATHER_API_KEY環境変数が設定されていません。天気情報は取得できません。")
	}
	
	// 各国の時間帯、都市名、OpenWeatherMapで使用する都市名
	locations := map[string]struct {
		DisplayName string
		CityForAPI  string
	}{
		"Asia/Tokyo":       {"東京 (日本)", "Tokyo"},
		"America/New_York": {"ニューヨーク (アメリカ)", "New York"},
		"Europe/London":    {"ロンドン (イギリス)", "London"},
		"Europe/Paris":     {"パリ (フランス)", "Paris"},
		"Australia/Sydney": {"シドニー (オーストラリア)", "Sydney"},
		"Asia/Shanghai":    {"上海 (中国)", "Shanghai"},
		"Asia/Singapore":   {"シンガポール", "Singapore"},
		"Europe/Moscow":    {"モスクワ (ロシア)", "Moscow"},
		"Asia/Dubai":       {"ドバイ (UAE)", "Dubai"},
		"America/Sao_Paulo": {"サンパウロ (ブラジル)", "Sao Paulo"},
	}
	
	// 各場所の時間と天気を表示
	for location, info := range locations {
		// 時間帯を取得
		loc, err := time.LoadLocation(location)
		if err != nil {
			fmt.Printf("時間帯の読み込みエラー %s: %v\n", location, err)
			continue
		}
		
		// 現地時間
		localTime := time.Now().In(loc)
		
		// 天気情報を取得
		weather := "情報なし"
		temperature := 0.0
		
		if apiKey != "" {
			weatherInfo, err := getWeather(info.CityForAPI, apiKey)
			if err != nil {
				fmt.Printf("天気情報の取得エラー %s: %v\n", info.CityForAPI, err)
			} else {
				if len(weatherInfo.Weather) > 0 {
					weather = translateWeather(weatherInfo.Weather[0].Main)
				}
				temperature = weatherInfo.Main.Temp
			}
		}
		
		// 表示
		fmt.Printf("%s: %s, 天気: %s, 気温: %.1f°C\n", 
			info.DisplayName, 
			localTime.Format("2006-01-02 15:04:05"),
			weather,
			temperature)
	}
}

// OpenWeatherMap APIから天気情報を取得
func getWeather(city, apiKey string) (WeatherResponse, error) {
	var weatherData WeatherResponse
	
	// スペースをURL用に変換
	city = strings.ReplaceAll(city, " ", "%20")
	
	// APIリクエストURL
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	
	// HTTPリクエスト
	resp, err := http.Get(url)
	if err != nil {
		return weatherData, err
	}
	defer resp.Body.Close()
	
	// レスポンスボディを読み取り
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return weatherData, err
	}
	
	// ステータスコードチェック
	if resp.StatusCode != http.StatusOK {
		return weatherData, fmt.Errorf("API error: %s", string(body))
	}
	
	// JSONをデコード
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return weatherData, err
	}
	
	return weatherData, nil
}
