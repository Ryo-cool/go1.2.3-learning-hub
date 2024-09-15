package main

import (
	"fmt"
)

// ジェネリック型エイリアス
type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) Contains(value T) bool {
	_, exists := s[value]
	return exists
}

func main() {
	fmt.Println("Generic type aliases:")

	intSet := NewSet[int]()
	intSet.Add(1)
	intSet.Add(2)
	intSet.Add(3)

	fmt.Printf("Set contains 2: %v\n", intSet.Contains(2))
	fmt.Printf("Set contains 4: %v\n", intSet.Contains(4))

	stringSet := NewSet[string]()
	stringSet.Add("apple")
	stringSet.Add("banana")

	fmt.Printf("Set contains 'apple': %v\n", stringSet.Contains("apple"))
	fmt.Printf("Set contains 'orange': %v\n", stringSet.Contains("orange"))
}
