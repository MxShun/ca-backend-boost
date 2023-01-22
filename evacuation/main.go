package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	DataSourceName = "postgres://admin:password@host.docker.internal:5432/pgdb"
)

type EvacuationArea struct {
	Name       string
	Address    string
	Flood      bool
	Landslide  bool
	Surge      bool
	Earthquake bool
	Tsunami    bool
	Fire       bool
	Inundation bool
	Volcano    bool
}

type ResponseJson struct {
	Name       string `json:"避難場所名"`
	Address    string `json:"所在地住所"`
	Flood      string `json:"洪水"`
	Landslide  string `json:"崖崩れ"`
	Surge      string `json:"高潮"`
	Earthquake string `json:"地震"`
	Tsunami    string `json:"津波"`
	Fire       string `json:"大規模な家事"`
	Inundation string `json:"内部氾濫"`
	Volcano    string `json:"火山活動"`
}

func prettifyBool(b bool) string {
	if b {
		return "◯"
	}
	return "✗"
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	requestName := request.QueryStringParameters["name"]
	hasRequestName := len(requestName) > 0
	requestAddress := request.QueryStringParameters["address"]
	hasRequestAddress := len(requestAddress) > 0

	db, err := sql.Open("pgx", DataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select * from evacuation_area")
	if err != nil {
		log.Fatal(err)
	}

	var response []ResponseJson
	for rows.Next() {
		row := EvacuationArea{}
		if err := rows.Scan(&row.Name, &row.Address, &row.Flood, &row.Landslide, &row.Surge, &row.Earthquake, &row.Tsunami, &row.Fire, &row.Inundation, &row.Volcano); err != nil {
			log.Fatal(err)
		}

		if (hasRequestName && strings.Contains(row.Name, requestName)) ||
			(hasRequestAddress && strings.Contains(row.Address, requestAddress)) {
			response = append(response, ResponseJson{
				row.Name,
				row.Address,
				prettifyBool(row.Flood),
				prettifyBool(row.Landslide),
				prettifyBool(row.Surge),
				prettifyBool(row.Earthquake),
				prettifyBool(row.Tsunami),
				prettifyBool(row.Fire),
				prettifyBool(row.Inundation),
				prettifyBool(row.Volcano),
			})
		}
	}

	var responseBytes, _ = json.Marshal(response)
	return events.APIGatewayProxyResponse{
		Body:       string(responseBytes),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
