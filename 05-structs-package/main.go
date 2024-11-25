package main

import (
	"fmt"
	"reflect"
	"structs"
)

type Person struct {
	Name string
	Age  int
	_    structs.HostLayout
}

func main() {
	fmt.Println("Go 1.23 構造体の例:")

	p := Person{Name: "Alice", Age: 30}
	v := reflect.ValueOf(p)

	fmt.Println("\nフィールドの列挙（新しいSeqメソッドを使用）:")
	for v := range v.Seq() {
		field := v.Interface()
		fmt.Printf("  %v\n", field)
	}

	m := make(map[string]interface{})
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if t.Field(i).Name != "_" {
			m[t.Field(i).Name] = v.Field(i).Interface()
		}
	}
	fmt.Println("\nMap:", m)
}
