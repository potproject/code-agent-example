package main

import (
	"fmt"
	"time"
)

// 時間を表示する簡単なgo言語のプログラム
func main() {
	now := time.Now()
	fmt.Println("The time is", now)
	
	year := now.Year()
	isLeapYear := year%4 == 0 && (year%100 != 0 || year%400 == 0)
	
	if isLeapYear {
		fmt.Printf("%d年はうるう年です\n", year)
	} else {
		fmt.Printf("%d年はうるう年ではありません\n", year)
	}
}
