package main

import (
	"database/sql"
	"encoding/csv"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	updateFileUrl  = "https://www.opendata.metro.tokyo.lg.jp/soumu/R4/130001_evacuation_area.csv"
	localFilePath  = "evacuation_area.csv"
	dataSourceName = "postgres://admin:password@localhost:5432/pgdb"
)

func shouldUpdate() bool {
	resp, err := http.Get(updateFileUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	lastModified, err := time.Parse(time.RFC1123, resp.Header.Get("Last-Modified"))
	if err != nil {
		log.Fatal(err)
	}

	isUpdatedToday := time.Since(lastModified) < 24*time.Hour
	if isUpdatedToday {
		return true
	}
	return false
}

func downloadFile() {
	resp, _ := http.Get(updateFileUrl)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	out, _ := os.Create(localFilePath)
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(out)

	_, err := io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func loadFile() [][]string {
	f, err := os.Open(localFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	r := csv.NewReader(transform.NewReader(f, unicode.UTF8.NewDecoder()))

	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func toBool(str string) bool {
	if str == "1" {
		return true
	}
	return false
}

func update(rows [][]string) (counter int) {
	db, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	dlt, err := db.Prepare("delete from evacuation_area;")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := dlt.Exec(); err != nil {
		log.Fatal(err)
	}

	for _, row := range rows {
		ist, err := db.Prepare("insert into evacuation_area (name, address, flood, landslide, surge, earthquake, tsunami, fire, inundation, volcano) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);")
		if err != nil {
			log.Fatal(err)
		}
		if _, err := ist.Exec(row[0], row[4], toBool(row[7]), toBool(row[8]), toBool(row[9]), toBool(row[10]), toBool(row[11]), toBool(row[12]), toBool(row[13]), toBool(row[14])); err != nil {
			log.Fatal(err)
		}

		counter++
	}

	return counter
}

func main() {
	if !shouldUpdate() {
		log.Println("no updates")
		return
	}
	downloadFile()
	data := loadFile()
	count := update(data)
	log.Printf("%d records updated", count)
}
