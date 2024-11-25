# Generic Type Aliases Example

このディレクトリでは、Go のジェネリック型エイリアスを使用した Set データ構造の実装例を示しています。

## 主な機能

1. `Set[T comparable]`: 任意の比較可能な型に対応する Set データ構造
2. `NewSet[T comparable]()`: 新しい Set を作成するコンストラクタ関数
3. `Add(value T)`: Set に要素を追加するメソッド
4. `Contains(value T)`: 要素の存在確認メソッド

## 実装の特徴

- ジェネリクスを使用した型安全な実装
- マップを使用した効率的な Set 実装
- `comparable`制約を使用した型制限

## 使用例

```go
// 整数のSet
intSet := NewSet[int]()
intSet.Add(1)
intSet.Add(2)

// 文字列のSet
stringSet := NewSet[string]()
stringSet.Add("apple")
stringSet.Add("banana")
```
