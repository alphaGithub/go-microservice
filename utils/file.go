package utils

import (
	"os"
)

func ReadFromFile(fileName string) ([]byte, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return content, nil
}
