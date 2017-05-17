package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"flag"
	"os"
)

type Location struct {
	Id string `json:"id"`
	Timestamp  string `json:"timestamp"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Type string `json:"type"`
}

func main() {
	var inputFile, outputFile string
	flag.StringVar(&inputFile, "f", "", "input csv file")
	flag.StringVar(&outputFile, "o", "result", "output json file")
	flag.Parse()

	csvFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord Location
	var allRecords []Location

	for i, each := range csvData {
		if i == 0 {
			continue
		}
		oneRecord.Id = each[0]
		oneRecord.Timestamp = each[1]
		oneRecord.Latitude = each[2]
		oneRecord.Longitude = each[3]
		oneRecord.Type = each[4]
		allRecords = append(allRecords, oneRecord)
	}

	jsonData, err := json.Marshal(allRecords)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))

	jsonFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
}
