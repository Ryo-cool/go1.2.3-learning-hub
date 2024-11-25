# Unique Package Example

このディレクトリでは、`golang.org/x/exp/slices`パッケージを使用して、スライスから重複要素を削除する方法を示しています。

## 主な機能

1. `slices.Compact`: スライスから連続する重複要素を削除する関数の使用例

## 実装の特徴

- `golang.org/x/exp/slices`パッケージの活用
- 異なる型（整数、文字列）での重複削除
- シンプルで効率的な実装

## 使用例

```go
// 整数スライスの重複削除
numbers := []int{1, 2, 2, 3, 4, 4, 5}
uniqueNumbers := slices.Compact(numbers)

// 文字列スライスの重複削除
fruits := []string{"apple", "banana", "apple", "cherry", "banana"}
uniqueFruits := slices.Compact(fruits)
```
