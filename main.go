package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body := []byte("<html><head><title>Go web service</title></head><body>HTMLのトップページです</body></html>")
		w.Write(body)
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		body := []byte("<html><head><title>Go web service</title></head><body>Pong</body></html>")
		w.Write(body)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
