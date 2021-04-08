package services

import (
	"github.com/Rob-a21/enjoy_recipe_backend/GoLang-Backend-/api/entity"
	recipe "github.com/Rob-a21/enjoy_recipe_backend/GoLang-Backend-/api/recipe"
)

type RecipeService struct {
	recipeRepo recipe.RecipeRepository
}

func NewRecipeService(recipeRepo recipe.RecipeRepository) *RecipeService {
	return &RecipeService{recipeRepo: recipeRepo}
}

func (cs *RecipeService) Recipies() ([]entity.Recipe, []error) {

	parties, errs := cs.recipeRepo.Recipies()
	if len(errs) > 0 {
		return nil, errs
	}
	return parties, errs

}

func (cs *RecipeService) Recipe(id uint32) (*entity.Recipe, []error) {
	par, errs := cs.recipeRepo.Recipe(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return par, errs

}

func (cs *RecipeService) UpdateRecipe(recipe *entity.Recipe) (*entity.Recipe, []error) {
	par, errs := cs.recipeRepo.UpdateRecipe(recipe)
	if len(errs) > 0 {
		return nil, errs
	}
	return par, errs

}

func (cs *RecipeService) DeleteRecipe(id uint32) (*entity.Recipe, []error) {

	par, errs := cs.recipeRepo.DeleteRecipe(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return par, errs
}

func (cs *RecipeService) StoreRecipe(recipe *entity.Recipe) (*entity.Recipe, []error) {

	par, errs := cs.recipeRepo.StoreRecipe(recipe)
	if len(errs) > 0 {
		return nil, errs
	}
	return par, errs
}
