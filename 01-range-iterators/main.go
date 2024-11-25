package main

import (
	"fmt"
	"iter" // Go 1.23の新しいパッケージ
)

// フィボナッチ数列のイテレータ
func fibonacci(limit int) iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for a < limit {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

// 偶数のみを返すイテレータ
func evenNumbers(limit int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for i := 0; i <= limit; i++ {
			if i%2 == 0 {
				if !yield(i, i*2) {
					return
				}
			}
		}
	}
}

func main() {
	fmt.Println("Go 1.23 イテレータの例:")

	// フィボナッチ数列の使用
	fmt.Println("\nフィボナッチ数列:")
	for v := range fibonacci(100) {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// インデックスと値のイテレーション
	fmt.Println("\n偶数とその2倍の値:")
	for i, v := range evenNumbers(10) {
		fmt.Printf("(%d: %d) ", i, v)
	}
	fmt.Println()
}
