// Программа, которая читаем в формате JSON и возвращает их в XML и наоборот.
package main

import (
	"day01/internal/dbcompare"
	"day01/internal/dbreader"
	"day01/pkg/utils"
)

func main() {
	//Получаем путь к файлам из аргументов командной строки
	filepathForOld, filepathForNew, err := utils.ParseFileForCompareDB()
	utils.HandleError(err)

	// Читаем файлы, десериализуем в структуры
	recipeOld, _, err := dbreader.ReadDB(filepathForOld)
	utils.HandleError(err)

	recipeNew, _, err := dbreader.ReadDB(filepathForNew)
	utils.HandleError(err)

	dbcompare.CompareDB(recipeOld, recipeNew)
}
