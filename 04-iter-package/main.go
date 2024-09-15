package main

import (
	"fmt"

	"golang.org/x/exp/iter"
)

func main() {
	fmt.Println("iter package example:")

	numbers := []int{1, 2, 3, 4, 5}

	// スライスの要素を2倍にする
	doubled := iter.Map(numbers, func(x int) int {
		return x * 2
	})

	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Doubled: %v\n", doubled)

	// 偶数のみをフィルタリング
	evens := iter.Filter(numbers, func(x int) bool {
		return x%2 == 0
	})

	fmt.Printf("Evens: %v\n", evens)
}
