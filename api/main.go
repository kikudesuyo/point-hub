package main

import (
	"fmt"
	"log"
	"os"

	"hoge/external"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".envの読み込みに失敗:", err)
	}

	cardNo := os.Getenv("TOEI_USER_ID")
	password := os.Getenv("TOEI_METRO_PASSWORD")
	if cardNo == "" || password == "" {
		log.Fatal("TOEI_USER_ID / TOEI_METRO_PASSWORD が未設定です")
	}

	client, err := external.NewToeiMetroClient()
	if err != nil {
		log.Fatal("クライアント初期化失敗:", err)
	}

	fmt.Println("ログイン中...")
	if err := client.Login(cardNo, password); err != nil {
		log.Fatal("ログイン失敗:", err)
	}
	fmt.Println("✅ ログイン成功")

	fmt.Println("データ取得中...")
	data, err := client.FetchAll()
	if err != nil {
		log.Fatal("データ取得失敗:", err)
	}

	fmt.Printf("\n── 会員情報 ──────────────────────\n")
	fmt.Printf("名前          : %s\n", data.Name)
	fmt.Printf("\n── ポイント ──────────────────────\n")
	fmt.Printf("保有ポイント  : %d pt\n", data.Point)
}
