package dbcompare

import (
	"day01/internal/dbreader"
	"fmt"
)

func CompareDB(recipeOld *dbreader.Recipes, recipeNew *dbreader.Recipes) {
	mapCakesOld := createRecipeMap(recipeOld)
	mapCakesNew := createRecipeMap(recipeNew)

	for _, oldRecipe := range mapCakesOld {
		// compare Name
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

func createRecipeMap(recipes *dbreader.Recipes) map[string]dbreader.Recipe {
	recipeMap := make(map[string]dbreader.Recipe)
	for _, recipe := range recipes.Cakes {
		recipeMap[recipe.Name] = recipe
	}
	return recipeMap
}

func compareTime(oldRecipe dbreader.Recipe, newRecipe dbreader.Recipe) {
	if oldRecipe.Time != newRecipe.Time {
		fmt.Printf("ИЗМЕНИЛОСЬ время готовки для торта %q - %q вместо %q\n",
			oldRecipe.Name, newRecipe.Time, oldRecipe.Time)
	}
}

func compareIngridients(oldRecipe dbreader.Recipe, newRecipe dbreader.Recipe) {
	mapIngredientsOld := createIngredientMap(oldRecipe.Ingredients)
	mapIngredientsNew := createIngredientMap(newRecipe.Ingredients)

	for _, ingridientOld := range mapIngredientsOld {
		if ingredientNew, exists := mapIngredientsNew[ingridientOld.Name]; !exists {
			fmt.Printf("УДАЛЕН ингредиент %q для торта %q\n", ingridientOld.Name, oldRecipe.Name)
		} else {
			compareIngredientDetails(oldRecipe.Name, ingridientOld, ingredientNew)
		}
	}
}

func createIngredientMap(ingredients []dbreader.Ingredient) map[string]dbreader.Ingredient {
	ingredientMap := make(map[string]dbreader.Ingredient)
	for _, ingredient := range ingredients {
		ingredientMap[ingredient.Name] = ingredient
	}
	return ingredientMap
}

func compareIngredientDetails(recipeName string, oldIngredient, newIngredient dbreader.Ingredient) {
	if newIngredient.Unit == "" && oldIngredient.Unit != "" {
		fmt.Printf("УДАЛЕНА единица измерения %q для ингредиента %q для торта %q\n",
			oldIngredient.Unit, oldIngredient.Name, recipeName)
	} else if oldIngredient.Unit == "" && newIngredient.Unit != "" {
		fmt.Printf("ДОБАВЛЕНА единица измерения %q для ингредиента %q для торта %q\n",
			newIngredient.Unit, oldIngredient.Name, recipeName)
	} else if oldIngredient.Unit != newIngredient.Unit {
		fmt.Printf("ИЗМЕНИЛАСЬ единица измерения для ингредиента %q для торта %q - %q вместо %q\n",
			oldIngredient.Name, recipeName, newIngredient.Unit, oldIngredient.Unit)
	}

	if newIngredient.Count != oldIngredient.Count {
		fmt.Printf("ИЗМЕНИЛОСЬ количество для ингредиента %q для торта %q - %q вместо %q\n",
			oldIngredient.Name, recipeName, newIngredient.Count, oldIngredient.Count)
	}
}
