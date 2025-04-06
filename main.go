package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 世界各国の時計と天気を表示するgo言語のプログラム
func main() {
	fmt.Println("=== 世界各国の時計と天気 ===")
	
	// 各国の時間帯と都市名
	locations := map[string]string{
		"Asia/Tokyo":       "東京 (日本)",
		"America/New_York": "ニューヨーク (アメリカ)",
		"Europe/London":    "ロンドン (イギリス)",
		"Europe/Paris":     "パリ (フランス)",
		"Australia/Sydney": "シドニー (オーストラリア)",
		"Asia/Shanghai":    "上海 (中国)",
		"Asia/Singapore":   "シンガポール",
		"Europe/Moscow":    "モスクワ (ロシア)",
		"Asia/Dubai":       "ドバイ (UAE)",
		"America/Sao_Paulo": "サンパウロ (ブラジル)",
	}

	// 天気の種類
	weatherTypes := []string{"晴れ", "曇り", "雨", "雪", "霧", "強風", "嵐"}
	
	// 気温の範囲（摂氏）
	minTemp := -10
	maxTemp := 40
	
	// Go 1.20以降ではrand.Seedは不要
	// 古いバージョンの場合のみこれを使用
	// rand.Seed(time.Now().UnixNano())
	
	// 各場所の時間と天気を表示
	for location, cityName := range locations {
		// 時間帯を取得
		loc, err := time.LoadLocation(location)
		if err != nil {
			fmt.Printf("時間帯の読み込みエラー %s: %v\n", location, err)
			continue
		}
		
		// 現地時間
		localTime := time.Now().In(loc)
		
		// ランダムな天気と気温を生成
		weather := weatherTypes[rand.Intn(len(weatherTypes))]
		temperature := rand.Intn(maxTemp-minTemp+1) + minTemp
		
		// 表示
		fmt.Printf("%s: %s, 天気: %s, 気温: %d°C\n", 
			cityName, 
			localTime.Format("2006-01-02 15:04:05"),
			weather,
			temperature)
	}
}
