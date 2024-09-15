package main

import (
	"fmt"

	"golang.org/x/exp/structs"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("structs package example:")

	p := Person{Name: "Alice", Age: 30}

	// 構造体のフィールドを列挙
	fields := structs.Fields(p)
	fmt.Println("Fields:")
	for _, field := range fields {
		fmt.Printf("  %s: %v\n", field.Name(), field.Value())
	}

	// 構造体をマップに変換
	m := structs.Map(p)
	fmt.Println("Map:", m)

	// 構造体のフィールド名を取得
	names := structs.Names(p)
	fmt.Println("Field names:", names)
}
