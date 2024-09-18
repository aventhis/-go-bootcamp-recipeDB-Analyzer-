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

	// mapCakesOld - мапа ключ - название торта / значение - структура Recipe
	mapCakesOld := make(map[string]dbreader.Recipe)
	// записали значения
	for _, structRecipe := range recipeOld.Cakes {
		mapCakesOld[structRecipe.Name] = structRecipe
	}

	mapCakesNew := make(map[string]dbreader.Recipe)
	for _, structRecipe := range recipeNew.Cakes {
		mapCakesNew[structRecipe.Name] = structRecipe
	}

	// recipesOld - структура Recipe (старая бд)
	// recipesNew - структура Recipe (новая бд)
	for _, recipesOld := range mapCakesOld {
		//сравниваем поля Name структуры Recipe
		if recipesNew, exists := mapCakesNew[recipesOld.Name]; !exists {
			fmt.Printf("REMOVED cake %v\n", recipesOld.Name)
		} else {
			// сравниваем поле Time структуры Recipe
			if recipesOld.Time != recipesNew.Time {
				fmt.Printf("CHANGED cooking time for cake %v - %v instead of %v\n", recipesOld.Name, recipesNew.Time, recipesOld.Time)
			}

			// делаем map - ключ - название ингредиента / значение - структура Ingredient
			mapIngredientsOld := make(map[string]dbreader.Ingredient)
			for _, ingredientOld := range recipesOld.Ingredients {
				mapIngredientsOld[ingredientOld.Name] = ingredientOld
			}

			mapIngredientsNew := make(map[string]dbreader.Ingredient)
			for _, ingredientNew := range recipesNew.Ingredients {
				mapIngredientsNew[ingredientNew.Name] = ingredientNew
			}

			for _, ingridientOld := range mapIngredientsOld {
				if ingridienNew, exists := mapIngredientsNew[ingridientOld.Name]; !exists {
					fmt.Printf("REMOVED ingredient %v for cake %v\n", ingridientOld.Name, recipesOld.Name)
				} else {
					fmt.Println(ingridienNew)
				}
			}

		}

	}

	for nameNew := range mapCakesNew {
		if _, exists := mapCakesOld[nameNew]; !exists {
			fmt.Printf("ADDED cake %v\n", nameNew)
		}
	}
}
