package importer

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

type CsvUser struct {
	Name  string
	GemId string
}

func Read(csvPath string) []CsvUser {
	csvLines := getCsvLines(csvPath)

	usersToAdd := []CsvUser{}
	gemidIdx, nameIdx := getCsvIdx(csvLines[0])

	for _, line := range csvLines[1:] {
		if line[gemidIdx] == "" {
			log.Default().Println("getCsvIdx - Warning! Possible error while parsing csv user ", line)
			continue
		}

		name := line[nameIdx]
		if name == "" {
			name = line[gemidIdx]
		}

		usersToAdd = append(usersToAdd, CsvUser{Name: name, GemId: line[gemidIdx]})
	}

	return usersToAdd
}

func getCsvLines(csvPath string) [][]string {
	csvFile, err := os.Open(csvPath)

	if err != nil {
		log.Fatalf("Error while opening csv file: %+v", err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	if strings.HasSuffix(csvPath, "tsv") {
		csvReader.Comma = '\t'
	} else if strings.HasSuffix(csvPath, "ssv") {
		csvReader.Comma = ';'
	}

	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Error while parsing csv: %+v", err)
	}

	if len(csvData) <= 1 {
		log.Fatal("Cannot import an empty CSV")
	}

	return csvData
}

func getCsvIdx(firstLine []string) (int, int) {
	nameIdx := -1
	gemidIdx := -1

	for i, field := range firstLine {
		if field == "name" {
			nameIdx = i
		} else if field == "gemid" {
			gemidIdx = i
		}
	}

	if gemidIdx == -1 {
		log.Fatal(`No column named "gemid" in the CSV file`)
	}

	if nameIdx == -1 {
		nameIdx = gemidIdx
	}

	return gemidIdx, nameIdx
}
