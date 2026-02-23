package converter

import (
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var fileValues = make([]float64, len(strings))
	for index, line := range strings {
		price, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}
		fileValues[index] = price
	}
	return fileValues, nil
}
