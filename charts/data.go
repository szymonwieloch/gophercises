package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func parseCSVFile(path string, columns []string) ([][]float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	header, err := csvReader.Read()
	if err != nil {
		return nil, err
	}
	colIndexes, err := columnsToIndexes(header, columns)
	if err != nil {
		return nil, err
	}
	return parseData(*csvReader, colIndexes)
}

func parseData(r csv.Reader, cols []int) ([][]float64, error) {
	result := make([][]float64, len(cols))
	for {
		rec, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		for idx, col := range cols {
			var val float64
			if rec[col] != "" {
				val, err = strconv.ParseFloat(rec[col], 64)
				if err != nil {
					return nil, fmt.Errorf("could not parse float value '%s' : %w", rec[col], err)
				}
			}

			result[idx] = append(result[idx], val)
		}
	}
	return result, nil
}

func columnsToIndexes(header, columns []string) ([]int, error) {
	var result []int
	for _, column := range columns {
		found := false
		for idx, name := range header {
			if name == column {
				result = append(result, idx)
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("Column '%s' not found", column)
		}
	}
	return result, nil
}
