package converter

import (
	"os"
	"strconv"
)

func StringsToFloats(data *os.File, strings []string) ([]float64, error) {
	var fileValues = make([]float64, len(strings))
	for index, line := range strings {
		price, err := strconv.ParseFloat(line, 64)
		if err != nil {
			data.Close()
			return nil, err
		}
		fileValues[index] = price
	}
	data.Close()
	return fileValues, nil
}
