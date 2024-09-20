package utils

import (
	"fmt"
	"os"
)

// ReadFile - Чтение файла
func ReadFile(filepath string) ([]byte, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %v", err)
	}
	return data, nil
}
