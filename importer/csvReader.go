package importer

import (
	"encoding/csv"
	"log"
	"os"
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
		log.Fatal("getCsvLines - Error while opening csv file: ", err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("getCsvLines - Error while parsing csv: ", err)
	}

	if len(csvData) <= 1 {
		log.Fatal("getCsvLines - Cannot import an empty CSV")
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
		log.Fatal(`getCsvIdx - No column named "gemid" in the CSV file`)
	}

	if nameIdx == -1 {
		nameIdx = gemidIdx
	}

	return gemidIdx, nameIdx
}
