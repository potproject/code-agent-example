package main

import (
	"fmt"
	"time"
)

// 時間を表示する簡単なgo言語のプログラム
func main() {
	now := time.Now()
	fmt.Println("現在の時刻は", now.Format("2006年01月02日 15時04分05秒 MST"))
}
