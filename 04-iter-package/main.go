package main

import (
	"fmt"
	"iter"
)

// スライスを2倍にするイテレータ
func double[T ~int | ~float64](s []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, x := range s {
			if !yield(x * 2) {
				return
			}
		}
	}
}

// 偶数のみをフィルタリングするイテレータ
func evens[T ~int](s []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, x := range s {
			if x%2 == 0 {
				if !yield(x) {
					return
				}
			}
		}
	}
}

func main() {
	fmt.Println("Go 1.23 iter パッケージの例:")

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original: %v\n", numbers)

	// スライスの要素を2倍にする
	fmt.Print("Doubled: ")
	for v := range double(numbers) {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// 偶数のみをフィルタリング
	fmt.Print("Evens: ")
	for v := range evens(numbers) {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
