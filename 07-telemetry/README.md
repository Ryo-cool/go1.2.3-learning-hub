# Telemetry Example

このディレクトリでは、Go のテレメトリ機能とビルド情報の取得方法を示しています。

## 主な機能

1. テレメトリのオプトイン設定
2. `debug.ReadBuildInfo`を使用したビルド情報の取得
3. モジュール依存関係の表示

## 実装の特徴

- 環境変数によるテレメトリ設定
- `runtime/debug`パッケージの活用
- モジュール情報の詳細な取得と表示

## 使用例

```go
// テレメトリの有効化
os.Setenv("GOTELEMETRY", "on")

// ビルド情報の取得
info, ok := debug.ReadBuildInfo()
if ok {
    // メインモジュール情報の表示
    fmt.Printf("メインモジュール: %s\n", info.Main.Path)

    // 依存モジュールの表示
    for _, dep := range info.Deps {
        fmt.Printf("%s: %s\n", dep.Path, dep.Version)
    }
}
```
