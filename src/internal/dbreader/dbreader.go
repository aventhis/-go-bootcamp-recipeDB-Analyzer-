package dbreader

import (
	"day01/pkg/utils"
	"errors"
	"path/filepath"
	"strings"
)

type Filetype string

const (
	JSON = "json"
	XML  = "xml"
)

// DBReader - интерфейс для чтения из файла JSON/XML и возврат структуры с данными
type DBReader interface {
	Read(data []byte) (*Recipes, error)
}

// NewDBReader - создает и возвращает реализацию интерфейса DBReader
func NewDBReader(filePath string) (DBReader, Filetype, error) {
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".json":
		return &JSONReader{}, JSON, nil
	case ".xml":
		return &XMLReader{}, XML, nil
	default:
		return nil, "", errors.New("неверный формат: файл должен быть расширения .json или .xml")
	}
}

func ReadDB(filepath string) (*Recipes, Filetype, error) {
	// Создаем ридер на основе типа файла (JSON или XML)
	reader, fileType, err := NewDBReader(filepath)
	if err != nil {
		return nil, "", err
	}

	// Читаем содержимое файла
	data, err := utils.ReadFile(filepath)
	if err != nil {
		return nil, "", err
	}

	// Десериализуем данные в структуру
	recipe, err := reader.Read(data)
	if err != nil {
		return nil, "", err
	}

	return recipe, fileType, nil
}
