package main

import (
	"encoding/csv"
	"os"
)

func readUsers(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var result []string
	for _, line := range lines {
		result = append(result, line[0])
	}
	return result, nil
}

func saveUsers(path string, users []string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0744)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	for _, user := range users {
		err = writer.Write([]string{user})
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return writer.Error()
}

func mergeUsers(existingUsers, newUsers []string) []string {
	result := make([]string, 0, len(existingUsers)+len(newUsers))
	known := make(map[string]bool, len(existingUsers)+len(newUsers))
	for _, eu := range existingUsers {
		known[eu] = true
		result = append(result, eu)
	}
	for _, nu := range newUsers {
		if !known[nu] {
			result = append(result, nu)
		}
	}
	return result
}
