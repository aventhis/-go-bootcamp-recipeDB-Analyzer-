package dbcompare

import (
	"day01/internal/dbreader"
	"fmt"
)

func CompareDB(recipeOld *dbreader.Recipes, recipeNew *dbreader.Recipes) {
	// mapCakesOld - мапа ключ - название торта / значение - структура Recipe
	mapCakesOld := make(map[string]dbreader.Recipe)
	for _, structRecipe := range recipeOld.Cakes {
		mapCakesOld[structRecipe.Name] = structRecipe
	}

	mapCakesNew := make(map[string]dbreader.Recipe)
	for _, structRecipe := range recipeNew.Cakes {
		mapCakesNew[structRecipe.Name] = structRecipe
	}

	// oldRecipe - структура Recipe (старая бд)
	// newRecipe - структура Recipe (новая бд)
	for _, oldRecipe := range mapCakesOld {
		//сравниваем поля Name структуры Recipe
		if newRecipe, exists := mapCakesNew[oldRecipe.Name]; !exists {
			fmt.Printf("УДАЛЕН торт %q\n", oldRecipe.Name)
		} else {
			compareTime(oldRecipe, newRecipe)
			compareIngridients(oldRecipe, newRecipe)
		}
	}

	for nameNew := range mapCakesNew {
		if _, exists := mapCakesOld[nameNew]; !exists {
			fmt.Printf("ДОБАВЛЕН торт %q\n", nameNew)
		}
	}
}

func compareTime(oldRecipe dbreader.Recipe, newRecipe dbreader.Recipe) {
	if oldRecipe.Time != newRecipe.Time {
		fmt.Printf("ИЗМЕНИЛОСЬ время готовки для торта %q - %q вместо %q\n",
			oldRecipe.Name, newRecipe.Time, oldRecipe.Time)
	}
}

func compareIngridients(oldRecipe dbreader.Recipe, newRecipe dbreader.Recipe) {
	// делаем map - ключ - название ингредиента / значение - структура Ingredient
	mapIngredientsOld := make(map[string]dbreader.Ingredient)
	for _, ingredientOld := range oldRecipe.Ingredients {
		mapIngredientsOld[ingredientOld.Name] = ingredientOld
	}

	mapIngredientsNew := make(map[string]dbreader.Ingredient)
	for _, ingredientNew := range newRecipe.Ingredients {
		mapIngredientsNew[ingredientNew.Name] = ingredientNew
	}

	//сравниваем поля Name структуры Ingridients
	for _, ingridientOld := range mapIngredientsOld {
		if ingridienNew, exists := mapIngredientsNew[ingridientOld.Name]; !exists {
			fmt.Printf("УДАЛЕН ингредиент %q для торта %q\n", ingridientOld.Name, oldRecipe.Name)
		} else {
			if ingridienNew.Unit == "" && ingridientOld.Unit != "" {
				fmt.Printf("УДАЛЕНА единица измерения %q для ингредиента %q для торта %q\n",
					ingridientOld.Unit, ingridientOld.Name, oldRecipe.Name)
			} else if ingridientOld.Unit == "" && ingridienNew.Unit != "" {
				fmt.Printf("ДОБАВЛЕНА единица измерения %q для ингредиента %q для торта %q\n",
					ingridientOld.Unit, ingridientOld.Name, oldRecipe.Name)
			} else if ingridientOld.Unit != ingridienNew.Unit {
				fmt.Printf("ИЗМЕНИЛАСЬ единица измерения для ингредиента %q для торта %q - %q "+
					"вместо %q\n", ingridientOld.Name, oldRecipe.Name, ingridientOld.Unit, ingridienNew.Unit)
			}
			if ingridienNew.Count != ingridientOld.Count {
				fmt.Printf("ИЗМЕНИЛОСЬ количество для ингредиента %q для торта %q - %q вместо %q\n",
					ingridientOld.Name, oldRecipe.Name, ingridienNew.Count, ingridientOld.Count)
			}
		}
	}
}
