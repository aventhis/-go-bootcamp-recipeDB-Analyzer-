// Программа, которая читаем в формате JSON и возвращает их в XML и наоборот.
package main

import (
	"day01/internal/dbreader"
	utils2 "day01/pkg/utils"
)

func main() {
	// Получаем путь к файлу из аргументов командной строки
	filepath, err := utils2.ParseFileForReadDB()
	utils2.HandleError(err)

	// Читаем файл и записываем данные в структуру
	recipe, fileType, err := dbreader.ReadDB(filepath)
	utils2.HandleError(err)

	// Преобразуем и выводим данные в противоположном формате
	err = dbreader.PrettyPainting(recipe, fileType)
	utils2.HandleError(err)
}
