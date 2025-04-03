package main

import (
	"fmt"
	"time"
)

// 時間を表示する簡単なgo言語のプログラム
func main() {
	now := time.Now()
	fmt.Println("L'heure actuelle est", now.Format("02 janvier 2006 15h04m05s MST"))
}
