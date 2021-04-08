package Recipe

import "github.com/Rob-a21/enjoy_recipe_backend/GoLang-Backend-/api/entity"

type RecipeServices interface {
	Recipies() ([]entity.Recipe, []error)
	Recipe(id uint32) (*entity.Recipe, []error)
	StoreRecipe(recipe *entity.Recipe) (*entity.Recipe, []error)
	UpdateRecipe(order *entity.Recipe) (*entity.Recipe, []error)
	DeleteRecipe(id uint32) (*entity.Recipe, []error)
}
