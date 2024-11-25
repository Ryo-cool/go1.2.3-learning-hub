# Time Changes Example

このディレクトリでは、Go の`time`パッケージを使用した時間関連の機能の実装例を示しています。

## 主な機能

1. `time.Timer`: 一度だけ発火するタイマーの使用
2. `time.Ticker`: 定期的に発火するティッカーの使用
3. `time.AfterFunc`: 指定時間後に関数を実行

## 実装の特徴

- 非同期タイマー処理
- チャネルを使用した時間制御
- ゴルーチンとの連携

## 使用例

```go
// タイマーの使用
timer := time.NewTimer(2 * time.Second)
<-timer.C

// ティッカーの使用
ticker := time.NewTicker(1 * time.Second)
<-ticker.C
ticker.Stop()

// AfterFuncの使用
time.AfterFunc(3*time.Second, func() {
    // 3秒後に実行される処理
})
```
