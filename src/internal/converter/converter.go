package converter

import (
	dbreader2 "day01/internal/dbreader"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
)

// PrettyPainting - Функция для конвертации и вывода данных в противоположном формате
func PrettyPainting(recipe *dbreader2.Recipes, fType dbreader2.Filetype) error {
	switch fType {
	case dbreader2.JSON:
		data, err := xml.MarshalIndent(recipe, "", "    ")
		if err != nil {
			return errors.New("ошибка при кодировании данных в JSON")
		}
		fmt.Println(string(data))
	case dbreader2.XML:
		data, err := json.MarshalIndent(recipe, "", "    ")
		if err != nil {
			return errors.New("ошибка при кодировании данных в XML")
		}
		fmt.Println(string(data))
	}
	return nil
}
