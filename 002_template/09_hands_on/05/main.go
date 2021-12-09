package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type record struct {
	Date     time.Time
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   int64
	AdjClose float64
}

func ReadCSV(fileName string) [][]string {
	originalFile, err := os.Open(fileName)
	defer originalFile.Close()
	if err != nil {
		log.Fatalln(err)
	}
	data, err := csv.NewReader(originalFile).ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func ProcessCSV(data [][]string) []record {
	var temp record

	result := make([]record, 0, len(data))

	for _, v := range data[1:] {
		temp.Date, _ = time.Parse("2006-01-02_multiplexer", v[0])
		temp.Open, _ = strconv.ParseFloat(v[1], 64)
		temp.High, _ = strconv.ParseFloat(v[2], 64)
		temp.Low, _ = strconv.ParseFloat(v[3], 64)
		temp.Close, _ = strconv.ParseFloat(v[4], 64)
		temp.Volume, _ = strconv.ParseInt(v[5], 10, 64)
		temp.AdjClose, _ = strconv.ParseFloat(v[6], 64)

		result = append(result, temp)
	}
	return result
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func handler(res http.ResponseWriter, req *http.Request) {
	// Parse CSV
	rawData := ReadCSV("table.csv")

	// Handle CSV
	processData := ProcessCSV(rawData)

	// Execute Template
	err := tpl.Execute(res, processData)
	if err != nil {
		log.Fatalln(err)
	}
}
