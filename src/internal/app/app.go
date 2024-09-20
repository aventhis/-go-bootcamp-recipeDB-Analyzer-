package app

import (
	"day01/internal/dbcompare"
	"day01/internal/dbreader"
	"day01/internal/utils"
)

// RunCompareDB запускает процесс чтения и сравнения баз данных
func RunCompareDB() {
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

// RunConvertDB запускает процесс чтения базы данных и преобразования форматов
func RunReadDB() {
	// Получаем путь к файлу из аргументов командной строки
	filepath, err := utils.ParseFileForReadDB()
	utils.HandleError(err)

	// Читаем файл и записываем данные в структуру
	recipe, fileType, err := dbreader.ReadDB(filepath)
	utils.HandleError(err)

	// Преобразуем и выводим данные в противоположном формате
	err = dbreader.PrettyPainting(recipe, fileType)
	utils.HandleError(err)
}
