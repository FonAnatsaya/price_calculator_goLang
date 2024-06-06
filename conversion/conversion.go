package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(stringVals []string) ([]float64, error) {

	var floatVals []float64
	for _, stringVal := range stringVals {
		floatVal, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}

		floatVals = append(floatVals, floatVal)
	}
	return floatVals, nil
}
