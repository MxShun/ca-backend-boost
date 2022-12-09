package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var days = flag.Int("days", 0, "日数の入力")
	var locale = flag.String("locale", "Asia/Tokyo", "ロケールの入力")

	// flag.Parse()

	location, err := time.LoadLocation(*locale)
	if err != nil {
			panic(err)
	}

	var now = time.Now().In(location)
	fmt.Println(now.AddDate(0, 0, *days))
}
