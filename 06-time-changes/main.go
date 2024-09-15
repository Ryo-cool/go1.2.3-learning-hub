package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time.Timer と time.Ticker の変更例:")

	// 新しい Timer の使用
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("タイマー開始")
	<-timer.C
	fmt.Println("タイマー終了")

	// 新しい Ticker の使用
	ticker := time.NewTicker(1 * time.Second)
	count := 0
	for {
		<-ticker.C
		count++
		fmt.Printf("ティック %d\n", count)
		if count >= 5 {
			ticker.Stop()
			break
		}
	}

	// AfterFunc の使用例
	done := make(chan bool)
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("AfterFunc: 3秒後に実行")
		done <- true
	})
	<-done
}
