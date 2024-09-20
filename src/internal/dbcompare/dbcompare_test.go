package dbcompare

import (
	"day01/internal/dbreader"
	"testing"
)

// Пример теста для идентичных баз данных
func TestCompareDB_NoChanges(t *testing.T) {
	recipeOld := &dbreader.Recipes{
		Cakes: []dbreader.Recipe{
			{
				Name: "Chocolate Cake",
				Time: "30 min",
				Ingredients: []dbreader.Ingredient{
					{Name: "Sugar", Count: "200g"},
					{Name: "Flour", Count: "300g"},
				},
			},
		},
	}

	recipeNew := &dbreader.Recipes{
		Cakes: []dbreader.Recipe{
			{
				Name: "Chocolate Cake",
				Time: "30 min",
				Ingredients: []dbreader.Ingredient{
					{Name: "Sugar", Count: "200g"},
					{Name: "Flour", Count: "300g"},
				},
			},
		},
	}

	CompareDB(recipeOld, recipeNew) // Ожидаем, что изменений не будет
}

// Пример теста, где удален один торт
func TestCompareDB_CakeRemoved(t *testing.T) {
	recipeOld := &dbreader.Recipes{
		Cakes: []dbreader.Recipe{
			{Name: "Vanilla Cake"},
			{Name: "Chocolate Cake"},
		},
	}

	recipeNew := &dbreader.Recipes{
		Cakes: []dbreader.Recipe{
			{Name: "Chocolate Cake"},
		},
	}

	CompareDB(recipeOld, recipeNew)
	// Ожидаем вывод: УДАЛЕН торт "Vanilla Cake"
}

// Пример теста, где изменилось время готовки
func TestCompareDB_TimeChanged(t *testing.T) {
	recipeOld := &dbreader.Recipes{
		Cakes: []dbreader.Recipe{
			{Name: "Chocolate Cake", Time: "30 min"},
		},
	}

	recipeNew := &dbreader.Recipes{
		Cakes: []dbreader.Recipe{
			{Name: "Chocolate Cake", Time: "45 min"},
		},
	}

	CompareDB(recipeOld, recipeNew)
	// Ожидаем вывод: ИЗМЕНИЛОСЬ время готовки для торта "Chocolate Cake" - "45 min" вместо "30 min"
}

// Пример теста, где изменилось количество ингредиентов
func TestCompareDB_IngredientChanged(t *testing.T) {
	recipeOld := &dbreader.Recipes{
		Cakes: []dbreader.Recipe{
			{
				Name: "Chocolate Cake",
				Ingredients: []dbreader.Ingredient{
					{Name: "Sugar", Count: "200g"},
				},
			},
		},
	}

	recipeNew := &dbreader.Recipes{
		Cakes: []dbreader.Recipe{
			{
				Name: "Chocolate Cake",
				Ingredients: []dbreader.Ingredient{
					{Name: "Sugar", Count: "250g"},
				},
			},
		},
	}

	CompareDB(recipeOld, recipeNew)
	// Ожидаем вывод: ИЗМЕНИЛОСЬ количество для ингредиента "Sugar" для торта "Chocolate Cake" - "250g" вместо "200g"
}
