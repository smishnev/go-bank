package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	balance, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("Failed reading file")
	}

	value, err := strconv.ParseFloat(string(balance), 64)

	if err != nil {
		return 1000, errors.New("Failed parsing file")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprintf("%.2f", value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
