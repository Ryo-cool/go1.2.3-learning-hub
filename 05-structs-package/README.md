# Structs Package Example

このディレクトリでは、`golang.org/x/exp/structs`パッケージを使用した構造体の操作方法を示しています。

## 主な機能

1. `structs.Fields`: 構造体のフィールド情報を取得
2. `structs.Map`: 構造体をマップに変換
3. `structs.Names`: 構造体のフィールド名を取得

## 実装の特徴

- リフレクションを使用した構造体の動的な操作
- 構造体からマップへの変換
- フィールド情報の取得と操作

## 使用例

```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 30}

// フィールド情報の取得
fields := structs.Fields(p)

// 構造体からマップへの変換
m := structs.Map(p)

// フィールド名の取得
names := structs.Names(p)
```
