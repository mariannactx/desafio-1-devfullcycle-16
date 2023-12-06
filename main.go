package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strings"
)

func readCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

func writeCsvFile(records [][]string) {
	file, err := os.Create(os.Args[2])
	defer file.Close()
	if err != nil {
			log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	var data [][]string
    for _, record := range records {
        data = append(data, record)
    }
    w.WriteAll(data)
}

func main() {
    records := readCsvFile(os.Args[1])

		// 		toLowerI := strings.ToLower(records[i][0])
		// 		byteArrayI := []byte(toLowerI)
		
		// 		toLowerJ := strings.ToLower(records[j][0])
		// 		byteArrayJ := []byte(toLowerJ)
		
		header := records[0:1]

		n := len(records)
		records = records[1 : n]
		
		sort.SliceStable(records, func(i, j int) bool {
			return records[i][1] < records[j][1]
		})

		sort.SliceStable(records, func(i, j int) bool {
			toLowerI := strings.ToLower(records[i][0])
			toLowerJ := strings.ToLower(records[j][0])
			
			return toLowerI < toLowerJ
		})
		
		newfile := header
		for i:=0; i < len(records); i++ {
			newfile = append(newfile, records[i])
		}

		writeCsvFile(newfile)
}