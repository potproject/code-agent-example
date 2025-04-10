package main

import (
	"fmt"
	"time"
)

// 時間を表示する簡単なgo言語のプログラム
func main() {
	now := time.Now()
	fmt.Printf("現在のUNIX時間は %d です。\n", now.UnixMilli())
}
