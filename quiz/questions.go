package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

// Parses CSV file with questions
func parseProblemsFile(path string) ([]problem, error) {

	result := []problem{}
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %v: %w", path, err)
	}
	csvReader := csv.NewReader(file)
	for {
		record, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to parse CSV file %v: %w", path, err)
		}
		if len(record) < 2 {
			return nil, fmt.Errorf("invalid number of elements in CSV row: %d", len(record))
		}
		result = append(result, problem{
			question: strings.TrimSpace(record[0]),
			answer:   strings.TrimSpace(record[1]),
		})
	}
	return result, nil
}
