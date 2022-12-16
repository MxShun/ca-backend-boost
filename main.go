package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestJson struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseJson struct {
	//デフォルトで変数名がキー名として設定される
	//{"Message":"JSONを返却","Status":200} の形式
	Message string
	Status  int

	//{"message":"JSONを返却","status":200} の形式
	//Message string `json:"message"`
	//Status  int    `json:"status"`
}

func main() {
	http.HandleFunc("/json", handleJson)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("サーバの起動に失敗")
	}
}

func handleJson(w http.ResponseWriter, r *http.Request) {
	var requestJson RequestJson
	json.NewDecoder(r.Body).Decode(&requestJson)

	fmt.Println(requestJson)

	responseJson := ResponseJson{"JSONを返却", 200}
	fmt.Println(responseJson)

	res, err := json.Marshal(responseJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

//curl -X POST http://localhost:8080/json -H 'Content-Type: application/json' -d '{"code":200, "message":"リクエストにJSONを設定"}'
