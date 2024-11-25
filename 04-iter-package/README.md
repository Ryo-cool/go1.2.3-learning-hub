# Iter Package Example

このディレクトリでは、`golang.org/x/exp/iter`パッケージを使用した関数型プログラミングスタイルのイテレーション処理を示しています。

## 主な機能

1. `iter.Map`: スライスの各要素に対して変換を適用
2. `iter.Filter`: 条件に基づいて要素をフィルタリング

## 実装の特徴

- 関数型プログラミングアプローチ
- 宣言的なデータ処理
- `golang.org/x/exp/iter`パッケージの活用

## 使用例

```go
numbers := []int{1, 2, 3, 4, 5}

// Map例：各要素を2倍に
doubled := iter.Map(numbers, func(x int) int {
    return x * 2
})

// Filter例：偶数のみを抽出
evens := iter.Filter(numbers, func(x int) bool {
    return x%2 == 0
})
```
