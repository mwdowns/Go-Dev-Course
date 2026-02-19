package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0.0, errors.New("no file found")
	}
	value, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0.0, errors.New("could not parse stored value")
	}
	return value, nil
}

func WriteFloatToFile(fileName string, value float64) {
	os.WriteFile(fileName, []byte(fmt.Sprintf("%f", value)), 0644)
}
