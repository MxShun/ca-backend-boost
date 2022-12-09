package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type payload struct {
	Text string `json:"text"`
}

func main() {
	postUrl := ""
	reqJson := payload{"5_ハロー世界."}
	bJson, err := json.Marshal(reqJson)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPost, postUrl, bytes.NewBuffer(bJson))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("%#v", string(byteArray))
}
