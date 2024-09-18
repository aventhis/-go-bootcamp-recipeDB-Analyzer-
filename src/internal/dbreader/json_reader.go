package dbreader

import (
	"encoding/json"
	"errors"
)

type JSONReader struct{}

// Функция читает из JSON и возвращает структуру
func (jr *JSONReader) Read(data []byte) (*Recipes, error) {
	var recipe Recipes
	if err := json.Unmarshal(data, &recipe); err != nil {
		return nil, errors.New("ошибка при декодировании JSON")
	}
	return &recipe, nil
}
