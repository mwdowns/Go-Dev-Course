package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadFile() (*os.File, []string, error) {
	data, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, nil, err
	}
	scanner := bufio.NewScanner(data)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		data.Close()
		return data, nil, err
	}
	return data, lines, nil
}

func (fm FileManager) WriteJson(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("failed to create file")
	}

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		file.Close()
		return errors.New("failed to encode file")
	}
	file.Close()
	return nil
}

func New(InputFilePath string, OutputFilePath string) FileManager {
	return FileManager{InputFilePath, OutputFilePath}
}
