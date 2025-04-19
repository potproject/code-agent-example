package main

import (
	"fmt"
	"time"
)

// 時間を表示する簡単なgo言語のプログラム
func main() {
	now := time.Now()
	fmt.Println("The time is", now)
	fmt.Println("7 days later:", now.AddDate(0, 0, 7))
	fmt.Println("10 days later:", now.AddDate(0, 0, 10))
	fmt.Println("14 days later:", now.AddDate(0, 0, 14))
}
