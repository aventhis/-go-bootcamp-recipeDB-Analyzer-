package fscompare

import (
	"bufio"
	"day01/internal/utils"
	"fmt"
)

func CompareFS(filepathForOld, filepathForNew string) {
	data := make(map[string]struct{})

	// Читаем старый файл
	fileOld, err := utils.OpenFile(filepathForOld)
	utils.HandleError(err)
	defer fileOld.Close()

	scanner := bufio.NewScanner(fileOld)
	for scanner.Scan() {
		txt := scanner.Text()
		data[txt] = struct{}{}
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("ошибка при чтении файла:", err)
		return
	}

	fileNew, err := utils.OpenFile(filepathForNew)
	utils.HandleError(err)
	defer fileNew.Close()

	// Читаем новый файл
	scanner = bufio.NewScanner(fileNew)
	for scanner.Scan() {
		txt := scanner.Text()
		if _, exist := data[txt]; !exist {
			fmt.Printf("ADDED %q\n", txt)
		} else {
			delete(data, txt)
		}
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	for path := range data {
		fmt.Printf("DELETE %q\n", path)
	}
}
