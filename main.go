package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	
	"github.com/joho/godotenv"
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

// 天気を取得して表示するプログラム
func main() {
	// コマンドライン引数の解析
	var locationName string
	flag.StringVar(&locationName, "location", "Tokyo", "天気を取得する都市名")
	
	// 場所選択
	locations := map[string]string{
		"tokyo":    "Tokyo",
		"la":       "Los Angeles",
		"osaka":    "Osaka",
		"hokkaido": "Sapporo",
	}
	
	locationsHelp := "利用可能な場所: "
	for k, v := range locations {
		locationsHelp += fmt.Sprintf("%s (%s) ", k, v)
	}
	
	location := flag.String("city", "", locationsHelp)
	flag.Parse()
	
	// 場所が指定されていれば、それを使用
	if *location != "" {
		if cityName, ok := locations[*location]; ok {
			locationName = cityName
		} else {
			fmt.Printf("未知の場所です: %s\n", *location)
			fmt.Println(locationsHelp)
			return
		}
	}
	
	// .envファイルから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .envファイルの読み込みに失敗しました。環境変数から直接APIキーを取得します。")
	}
	
	// 指定された場所の天気を取得
	weather, err := getWeather(locationName)
	if err != nil {
		fmt.Println("天気情報の取得に失敗しました:", err)
		return
	}

	// 日本のタイムゾーンを設定
	jst := time.FixedZone("JST", 9*60*60)
	now := time.Now().In(jst)
	
	// 結果を表示
	fmt.Printf("日時: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("場所: %s\n", weather.Name)
	if len(weather.Weather) > 0 {
		fmt.Printf("天気: %s\n", weather.Weather[0].Description)
	}
	fmt.Printf("気温: %.1f°C\n", weather.Main.Temp)
}

// 指定された場所の天気情報を取得する関数
func getWeather(location string) (*WeatherResponse, error) {
	// 環境変数からAPIキーを取得
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("APIキーが設定されていません。.envファイルまたは環境変数にOPENWEATHERMAP_API_KEYを設定してください")
	}
	
	// OpenWeatherMap APIのエンドポイント
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", location, apiKey)

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
