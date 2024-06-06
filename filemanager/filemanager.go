package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func Readlines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("reading line in file failed")
	}
	file.Close()
	return lines, nil
}

func WriteJSON(path string, data interface{}) error {
	// file, err := os.Create(path)

	// encoder := json.NewEncoder(file)
	// err = encoder.Encode(data)

	json, err := json.Marshal(data)

	if err != nil {
		return errors.New("failed to convert data to json")
	}

	os.WriteFile(path, json, 0644)
	return nil
}
