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
	isLeapYear := (year%4 == 0 && year%100 != 0) || (year%400 == 0)
	fmt.Printf("Year %d is a leap year: %t\n", year, isLeapYear)
}
