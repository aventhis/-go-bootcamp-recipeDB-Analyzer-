package dbreader

import (
	"encoding/xml"
	"errors"
)

type XMLReader struct{}

// Функция читает из XML и возвращает структуру
func (xr *XMLReader) Read(data []byte) (*Recipes, error) {
	var recipe Recipes
	if err := xml.Unmarshal(data, &recipe); err != nil {
		return nil, errors.New("ошибка при декодировании XML")
	}
	return &recipe, nil
}
