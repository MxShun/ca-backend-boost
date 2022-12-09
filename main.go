package main

import (
	"encoding/csv"
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("io/130001_public_facility.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	r := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))

	for {
		records, err := r.Read()
		if err != nil {
			log.Fatal(err)
		}

		address := records[8]
		if !strings.Contains(address, "東京都台東区") {
			continue
		}
		fmt.Println(records)
	}
}
