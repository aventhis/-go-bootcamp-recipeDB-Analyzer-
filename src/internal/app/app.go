package app

import (
	"bufio"
	"day01/internal/dbcompare"
	"day01/internal/dbreader"
	"day01/internal/utils"
	"fmt"
	"os"
)

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

func RunCompareFS() {
	filepathForOld, filepathForNew, err := utils.ParseFileForCompareDB()
	utils.HandleError(err)

	fileOld, err := os.Open(filepathForOld)
	if err != nil {
		fmt.Println("Ошибка открытия файла", err)
	}
	defer fileOld.Close()

	data := make(map[string]struct{})
	scanner := bufio.NewScanner(fileOld)
	for scanner.Scan() {
		txt := scanner.Text()
		data[txt] = struct{}{}
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	fileNew, err := os.Open(filepathForNew)
	if err != nil {
		fmt.Println("Ошибка открытия файла", err)
	}
	defer fileNew.Close()

	scanner = bufio.NewScanner(fileNew)
	for scanner.Scan() {
		txt := scanner.Text()
		if _, exist := data[txt]; !exist {
			fmt.Printf("ADDED %q\n", txt)
		} else {
			delete(data, txt)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	for path := range data {
		fmt.Printf("DELETE %q\n", path)
	}
}
