package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

func main() {
	fmt.Println("テレメトリ機能の例:")

	// テレメトリのオプトイン設定
	os.Setenv("GOTELEMETRY", "on")

	// ビルド情報の取得
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("ビルド情報を取得できませんでした")
		return
	}

	// モジュール情報の表示
	fmt.Println("モジュール情報:")
	fmt.Printf("  メインモジュール: %s\n", info.Main.Path)
	fmt.Printf("  バージョン: %s\n", info.Main.Version)

	// 依存モジュールの表示
	fmt.Println("依存モジュール:")
	for _, dep := range info.Deps {
		fmt.Printf("  - %s: %s\n", dep.Path, dep.Version)
	}

	// 注意: 実際のテレメトリデータの収集と送信は
	// Goランタイムによって内部的に処理されます
	fmt.Println("\nテレメトリ機能が有効化されました（シミュレーション）")
}
