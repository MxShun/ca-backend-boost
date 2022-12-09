package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://www.cyberagent.co.jp", nil)
	if err != nil {
		panic(err)
	}

	// リクエスト送信
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	// クローズ
	defer resp.Body.Close()

	// HTTP ステータスコード
	fmt.Println(resp.StatusCode)
	// ヘッダ情報
	fmt.Println(resp.Header)
	// レスポンスの長さ
	fmt.Println(resp.ContentLength)
	// ボディ部分の読み込み
	body, err := io.ReadAll(resp.Body)
	if (err != nil) {
		panic(err)
	}
	fmt.Println(body)
}
