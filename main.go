package main

import (
	"fmt"
	"time"
)

// 時間を表示する簡単なgo言語のプログラム
func main() {
	now := time.Now()
	fmt.Println("現在時刻:", now)
	fmt.Println("7日後:", now.AddDate(0, 0, 7))
	fmt.Println("10日後:", now.AddDate(0, 0, 10))
	fmt.Println("14日後:", now.AddDate(0, 0, 14))
	fmt.Println("30日後:", now.AddDate(0, 0, 30))
	fmt.Println("1年後:", now.AddDate(1, 0, 0))
}
