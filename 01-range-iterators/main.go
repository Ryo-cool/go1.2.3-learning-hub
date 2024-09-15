package main

import (
	"fmt"
)

// カスタムイテレータ関数
func countTo(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

// フィボナッチ数列のイテレータ
func fibonacci(limit int) <-chan int {
	ch := make(chan int)
	go func() {
		a, b := 0, 1
		for a < limit {
			ch <- a
			a, b = b, a+b
		}
		close(ch)
	}()
	return ch
}

func main() {
	fmt.Println("1. 新しいイテレータ関数の使用方法:")

	// 1から5までの数をイテレート
	for v := range countTo(5) {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// フィボナッチ数列を100未満までイテレート
	for v := range fibonacci(100) {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	fmt.Println("\n2. カスタムイテレータの作成と使用:")

	// 偶数のみを返すカスタムイテレータ
	evenNumbers := countTo(10)
	for v := range evenNumbers {
		if v%2 == 0 {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()

	fmt.Println("\n3. ジェネリック型エイリアスの使用:")

	for v := range evenNumbers {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
