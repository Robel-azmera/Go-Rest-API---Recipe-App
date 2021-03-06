package handlerpackage

import (
	"encoding/json"
	"fmt"

	"github.com/Rob-a21/enjoy_recipe_backend/GoLang-Backend-/api/entity"
	recipe "github.com/Rob-a21/enjoy_recipe_backend/GoLang-Backend-/api/recipe"

	"github.com/Rob-a21/flutter_backend/api/utils"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type RecipeApiHandler struct {
	recipeServices recipe.RecipeRepository
}

func NewRecipeApiHandler(recipeServices recipe.RecipeRepository) *RecipeApiHandler {
	return &RecipeApiHandler{recipeServices: recipeServices}
}
func (cah *RecipeApiHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	cl, errs := cah.recipeServices.Recipe(uint32(id))
	print(cl)
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(cl, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

func (cah *RecipeApiHandler) GetRecipies(w http.ResponseWriter, r *http.Request) {

	elections, errs := cah.recipeServices.Recipies()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	output, err := json.MarshalIndent(elections, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

func (cah *RecipeApiHandler) PostRecipe(w http.ResponseWriter, r *http.Request) {

	body := utils.BodyParser(r)
	var par entity.Recipe
	err := json.Unmarshal(body, &par)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	storeRecipe, errs := cah.recipeServices.StoreRecipe(&par)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return

	}
	output, err := json.MarshalIndent(storeRecipe, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}

func (fah *RecipeApiHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("125")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rcp, errs := fah.recipeServices.Recipe(uint32(id))
	var recipeWithStringTimeStamp RecipeWithStringTimeStamp
	if len(errs) > 0 {
		fmt.Println("134")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	body := utils.BodyParser(r)
	err = json.Unmarshal(body, &recipeWithStringTimeStamp)

	if err != nil {
		fmt.Println("153")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// rcp.ID = recipeWithStringTimeStamp.ID
	rcp.RecipeName = recipeWithStringTimeStamp.RecipeName
	rcp.Causions = recipeWithStringTimeStamp.Causions
	rcp.Calories = recipeWithStringTimeStamp.Calories
	rcp.Instructions = recipeWithStringTimeStamp.Instructions
	rcp.Image = recipeWithStringTimeStamp.Image

	rcp, errs = fah.recipeServices.UpdateRecipe(rcp)
	if len(errs) > 0 {
		fmt.Println("169")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(rcp, "", "\t\t")

	if err != nil {
		fmt.Println("170")
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}

func (cah *RecipeApiHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := cah.recipeServices.DeleteRecipe(uint32(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}

type RecipeWithStringTimeStamp struct {
	ID           int    `gorm:"primary_key;auto_increment" json:"id"`
	Image        string `gorm:"type:varchar(255);not null" json:"image"`
	RecipeName   string `gorm:"type:varchar(255);not null" json:"name"`
	Causions     string `gorm:"type:varchar(255);not null" json:"causions"`
	Instructions string `gorm:"type:varchar(255);not null" json:"instructions"`
	Calories     int    ` json:"calories"`
}
