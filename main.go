package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/input", handleQuery)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func handleQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())

	name := r.URL.Query().Get("name")
	price := r.URL.Query().Get("price")

	fmt.Println(name, price)
	fmt.Fprint(w, "OK")
}
