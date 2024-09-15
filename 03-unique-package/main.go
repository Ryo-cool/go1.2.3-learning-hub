package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	fmt.Println("unique package example:")

	// スライスから重複を削除
	numbers := []int{1, 2, 2, 3, 4, 4, 5}
	uniqueNumbers := slices.Compact(numbers)
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Unique: %v\n", uniqueNumbers)

	// 文字列スライスから重複を削除
	fruits := []string{"apple", "banana", "apple", "cherry", "banana"}
	uniqueFruits := slices.Compact(fruits)
	fmt.Printf("Original: %v\n", fruits)
	fmt.Printf("Unique: %v\n", uniqueFruits)
}
