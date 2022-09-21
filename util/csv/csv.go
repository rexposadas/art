package csv

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

// ParseCSV returns a slice of lines.
// A line is a list in the CSV.
// We don't return an error.  Panic whenever we encounter an error.
// This is OK since so we don't continue processing if we can't
// process the file.
func ParseCSV(filename string) [][]string {
	if filename == "" {
		log.Fatal("empty filename given to csv parser")
	}

	csvFile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to read file %q, %s", filename, err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))
	var list [][]string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error reading line from csv file %s", err)
		}

		list = append(list, line)
	}

	return list
}
