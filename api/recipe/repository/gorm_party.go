package repository

import (
	"github.com/Rob-a21/enjoy_recipe_backend/GoLang-Backend-/api/entity"
	"github.com/jinzhu/gorm"
)

type RecipeGormRepo struct {
	conn *gorm.DB
}

func NewRecipeGormRepo(db *gorm.DB) *RecipeGormRepo {
	return &RecipeGormRepo{conn: db}
}

func (recipeRepo *RecipeGormRepo) Recipies() ([]entity.Recipe, []error) {
	var recipes []entity.Recipe
	errs := recipeRepo.conn.Find(&recipes).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return recipes, errs
}

func (recipeRepo *RecipeGormRepo) Recipe(id uint32) (*entity.Recipe, []error) {
	recipe := entity.Recipe{}
	errs := recipeRepo.conn.First(&recipe, id).GetErrors()
	if len(errs) > 0 {
		//fmt.Println(errs)
		return nil, errs
	}
	return &recipe, errs
}

//func (partyRepo *PartyGormRepo) UpdateParty(party *entity.Party) (*entity.Party, []error) {
//	errs := partyRepo.conn.Save(party).GetErrors()
//	if len(errs) > 0 {
//		return nil, errs
//	}
//	return party, errs
//}
// =========================================================
func (recipeRepo *RecipeGormRepo) UpdateRecipe(party *entity.Recipe) (*entity.Recipe, []error) {
	errs := recipeRepo.conn.Save(party).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return party, errs
}

func (recipeRepo *RecipeGormRepo) DeleteRecipe(id uint32) (*entity.Recipe, []error) {
	party, errs := recipeRepo.Recipe(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = recipeRepo.conn.Delete(party, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return party, errs
}

func (recipeRepo *RecipeGormRepo) StoreRecipe(party *entity.Recipe) (*entity.Recipe, []error) {
	errs := recipeRepo.conn.Create(party).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return party, errs
}
