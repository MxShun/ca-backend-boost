package main

import (
	"encoding/csv"
	"encoding/json"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	UpdateFileUrl = "https://www.opendata.metro.tokyo.lg.jp/soumu/R4/130001_evacuation_area.csv"
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

// TODO: API が増えたら Swagger 等でドキュメント化する
func usage() string {
	return `
		Evacuation API Usage
		東京都の避難所・避難場所を検索する API です

		[GET] /evacuation{?key,flood,landslide,surge,earthquake,tsunami,fire,inundation,volcano}

		◯ Query parameters
		- key : 検索キーワード string (required)
		- flood : 洪水時避難所・避難場所 bool (optional)
		- landslide : 崖崩れ時避難所・避難場所 bool (optional)
		- surge : 高潮時避難所・避難場所 bool (optional)
		- earthquake : 地震時避難所・避難場所 bool (optional)
		- tsunami : 津波時避難所・避難場所 bool (optional)
		- fire : 大規模火災時避難所・避難場所 bool (optional)
		- inundation : 内水氾濫時避難所・避難場所 bool (optional)
		- volcano : 火山活動時避難所・避難場所 bool (optional)

		◯ Example
		/evacuation?key=池袋&earthquake&tsunami : 「池袋」を含む地震もしくは津波時避難所・避難場所を検索
		`
}

func toBool(str string) bool {
	if str == "1" {
		return true
	}
	return false
}

func prettifyBool(b bool) string {
	if b {
		return "◯"
	}
	return "✗"
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	key := request.QueryStringParameters["key"]
	hasKey := len(key) > 0

	_, hasFlood := request.QueryStringParameters["flood"]
	_, hasLandslide := request.QueryStringParameters["landslide"]
	_, hasSurge := request.QueryStringParameters["surge"]
	_, hasEarthquake := request.QueryStringParameters["earthquake"]
	_, hasTsunami := request.QueryStringParameters["tsunami"]
	_, hasFire := request.QueryStringParameters["fire"]
	_, hasInundation := request.QueryStringParameters["inundation"]
	_, hasVolcano := request.QueryStringParameters["volcano"]

	if !hasKey {
		return events.APIGatewayProxyResponse{
			Body:       usage(),
			StatusCode: 200,
		}, nil
	}

	// FIXME: いただいた DB 認証情報が不正だったので DB 接続ができない
	//db, err := sql.Open("pgx", DataSourceName)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//rows, err := db.Query("select * from evacuation_area")
	//if err != nil {
	//	log.Fatal(err)
	//}
	// なので、DB を使わず都度 CSV を読む
	resp, _ := http.Get(UpdateFileUrl)
	defer resp.Body.Close()
	body, _ := csv.NewReader(transform.NewReader(resp.Body, unicode.UTF8.NewDecoder())).ReadAll()
	var rows []EvacuationArea
	for _, row := range body {
		rows = append(rows, EvacuationArea{
			row[0],
			row[4],
			toBool(row[7]),
			toBool(row[8]),
			toBool(row[9]),
			toBool(row[10]),
			toBool(row[11]),
			toBool(row[12]),
			toBool(row[13]),
			toBool(row[14]),
		})
	}

	var response []ResponseJson
	for _, row := range rows {
		if !strings.Contains(row.Name, key) && !strings.Contains(row.Address, key) {
			continue
		}

		if hasFlood && !row.Flood {
			continue
		}
		if hasLandslide && !row.Landslide {
			continue
		}
		if hasSurge && !row.Surge {
			continue
		}
		if hasEarthquake && !row.Earthquake {
			continue
		}
		if hasTsunami && !row.Tsunami {
			continue
		}
		if hasFire && !row.Fire {
			continue
		}
		if hasInundation && !row.Inundation {
			continue
		}
		if hasVolcano && !row.Volcano {
			continue
		}

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

	if len(response) > 0 {
		responseBytes, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		return events.APIGatewayProxyResponse{
			Body:       string(responseBytes),
			StatusCode: 200,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       "検索条件に一致するデータがありません",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
