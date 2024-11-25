# Range Iterators Example

このディレクトリでは、Go でのカスタムイテレータの実装と使用方法を示しています。

## 主な機能

1. `countTo(n int)`: 1 から n までの数値を生成するイテレータ
2. `fibonacci(limit int)`: 指定された上限までのフィボナッチ数列を生成するイテレータ

## 実装の特徴

- チャネルを使用したイテレータの実装
- `range`構文を使用したイテレータの使用
- ゴルーチンを使用した非同期処理

## 使用例

```go
// 1から5までの数をイテレート
for v := range countTo(5) {
    fmt.Printf("%d ", v)
}

// フィボナッチ数列を100未満までイテレート
for v := range fibonacci(100) {
    fmt.Printf("%d ", v)
}
```
