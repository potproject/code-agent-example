package main

import (
	"fmt"
	"time"
)

// 世界各国の時間を表示するプログラム
func main() {
	now := time.Now().UTC()
	
	locations := map[string]string{
		"UTC":       "UTC",
		"Tokyo":     "Asia/Tokyo",
		"New York":  "America/New_York",
		"London":    "Europe/London",
		"Paris":     "Europe/Paris",
		"Sydney":    "Australia/Sydney",
		"Beijing":   "Asia/Shanghai",
		"Moscow":    "Europe/Moscow",
		"Dubai":     "Asia/Dubai",
		"Sao Paulo": "America/Sao_Paulo",
	}
	
	fmt.Println("Current time in various locations around the world:")
	fmt.Println("---------------------------------------------------")
	
	for name, loc := range locations {
		location, err := time.LoadLocation(loc)
		if err != nil {
			fmt.Printf("Error loading location %s: %v\n", loc, err)
			continue
		}
		
		localTime := now.In(location)
		fmt.Printf("%-10s: %s\n", name, localTime.Format("2006-01-02 15:04:05 MST"))
	}
}
