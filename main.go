package main

import (
	"log"
	"net/http"
	"text/template"
)

type TemplateIndex struct {
	Title string
	Body  string
}

func main() {
	http.HandleFunc("/", handlerIndex)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}

	templateIndex := TemplateIndex{"タイトル", "本文"}

	if err := t.Execute(w, templateIndex); err != nil {
		log.Fatalf("テンプレートの埋め込みエラー: %v", err)
	}
}
