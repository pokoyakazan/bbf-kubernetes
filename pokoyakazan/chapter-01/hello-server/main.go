package main

import (
	"fmt"
	// "github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// "/"で始まるパスへのリクエストを受け付けるハンドラ
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	// ログにサーバー起動のメッセージを出力
	log.Println("Starting server on port 8080")
	// ポート8080でHTTPサーバーを起動
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
