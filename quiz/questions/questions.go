package questions

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Question struct {
	Question string
	Answer   string
}

// Parses CSV file with questions
func ParseFile(path string) ([]Question, error) {

	result := []Question{}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	csvReader := csv.NewReader(file)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if len(record) < 2 {
			return nil, fmt.Errorf("invalid number of elements in CSV row: %d", len(record))
		}
		result = append(result, Question{
			Question: record[0],
			Answer:   record[1],
		})
	}
	return result, nil
}
