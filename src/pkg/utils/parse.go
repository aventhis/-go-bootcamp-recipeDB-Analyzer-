package utils

import (
	"flag"
	"fmt"
)

// ParseFileForReadDB - Парсинг аргументов командной строки (флага f + путь к файлу)
func ParseFileForReadDB() (string, error) {
	filepath := flag.String("f", "", "путь к файлу")
	flag.Parse()

	// проверяем на наличие флага f
	if *filepath == "" {
		return "", fmt.Errorf("ошибка: необходимо указать флаг -f и путь к файлу\nusage: ./readDB -f <filename>")
	}

	// проверяем, что передан один файл
	if flag.NArg() > 0 {
		return "", fmt.Errorf("Ошибка: можно указать только один файл\nusage: ./readDB -f <filename>")
	}
	return *filepath, nil
}

// ParseFileForCompareDB - Парсинг аргументов командной строки (флага old и new + пути к файлам)
func ParseFileForCompareDB() (string, string, error) {
	oldFlag := flag.String("old", "", "путь к файлу до оригинальной базы данных")
	newFlag := flag.String("new", "", "путь к файлу до новой базы данных")
	flag.Parse()

	// проверяем на наличие флага f
	if *oldFlag == "" || *newFlag == "" {
		return "", "", fmt.Errorf("ошибка: необходимо указать оба флага --old и --new\nusage: ~$ ./compareDB --old original_database.xml --new stolen_database.json")
	}

	// проверяем, что передан один файл
	if flag.NArg() > 0 {
		return "", "", fmt.Errorf("Ошибка: можно указать только один файл\nusage: ./readDB -f <filename>")
	}

	return *oldFlag, *newFlag, nil
}
