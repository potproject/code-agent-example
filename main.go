package main

import (
	"fmt"
	"time"
)

// 時間を表示する簡単なgo言語のプログラム
func main() {
	now := time.Now()
	fmt.Println("The time is", now)
	
	// 7日後の時刻
	future7 := now.AddDate(0, 0, 7)
	fmt.Println("7日後の時刻は", future7)
	
	// 10日後の時刻
	future10 := now.AddDate(0, 0, 10)
	fmt.Println("10日後の時刻は", future10)
	
	// 14日後の時刻
	future14 := now.AddDate(0, 0, 14)
	fmt.Println("14日後の時刻は", future14)
}
