// Программа, которая читаем в формате JSON и возвращает их в XML и наоборот.
package main

import (
	"day01/internal/dbreader"
	"day01/pkg/utils"
	"fmt"
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

	// mapCakesOld - мапа для названия торта и структуры Recipe
	mapCakesOld := make(map[string]dbreader.Recipe)

	// ключ, значение nameCake - содержит структуру Recipe
	for _, structRecipe := range recipeOld.Cakes {
		mapCakesOld[structRecipe.Name] = structRecipe
	}

	mapCakesNew := make(map[string]dbreader.Recipe)

	for _, structRecipe := range recipeNew.Cakes {
		mapCakesNew[structRecipe.Name] = structRecipe
	}

	for nameOld, cakeOld := range mapCakesOld {
		if cakeNew, exists := mapCakesNew[nameOld]; !exists {
			fmt.Printf("REMOVED cake %v\n", nameOld)
		} else {
			if cakeOld.Time != cakeNew.Time {
				fmt.Printf("CHANGED cooking time for cake %v - %v instead of %v\n", nameOld, cakeNew.Time, cakeOld.Time)
			}

			mapIngredientsOld := make(map[string]dbreader.Ingredient)
			for _, ingredientOld := range cakeOld.Ingredients {
				mapIngredientsOld[ingredientOld.Name] = ingredientOld
			}

			mapIngredientsNew := make(map[string]dbreader.Ingredient)
			for _, ingredientNew := range cakeOld.Ingredients {
				mapIngredientsNew[ingredientNew.Name] = ingredientNew
			}

			for ingrNameOld, _ := range mapIngredientsOld {
				if _, exists := mapIngredientsNew[ingrNameOld]; !exists {
					fmt.Printf("REMOVED ingredient %v for cake %v\n", ingrNameOld, nameOld)
				}
			}

		}

	}

	for nameNew, _ := range mapCakesNew {
		if _, exists := mapCakesOld[nameNew]; !exists {
			fmt.Printf("ADDED cake %v\n", nameNew)
		}
	}
}
