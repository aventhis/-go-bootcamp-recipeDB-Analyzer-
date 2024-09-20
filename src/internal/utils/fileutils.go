package utils

import (
	"errors"
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

func OpenFile(filepath string) (*os.File, error) {
	filePath, err := os.Open(filepath)
	if err != nil {
		return nil, errors.New("ошибка открытия файла")
	}
	return filePath, nil
}
