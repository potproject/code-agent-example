package main

import (
	"fmt"
	"time"
)

// 日本の干支を取得する関数
func getJapaneseZodiac(year int) string {
	// 干支の配列（子、丑、寅、卯、辰、巳、午、未、申、酉、戌、亥）
	zodiac := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	// 年から干支のインデックスを計算（2008年は子年なので、その差を計算して12で割った余りを使用）
	index := (year - 2008) % 12
	if index < 0 {
		index += 12
	}
	return zodiac[index]
}

// 時間を表示する簡単なgo言語のプログラム
func main() {
	now := time.Now()
	fmt.Println("The time is", now)
	
	year := now.Year()
	isLeapYear := (year%4 == 0 && year%100 != 0) || (year%400 == 0)
	fmt.Printf("Year %d is a leap year: %t\n", year, isLeapYear)
	
	// 干支を表示
	zodiac := getJapaneseZodiac(year)
	fmt.Printf("Year %d is the year of %s (干支: %s年)\n", year, zodiac, zodiac)
}
