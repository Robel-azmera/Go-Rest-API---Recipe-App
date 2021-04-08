package main

import (
	"fmt"

	"github.com/gorilla/mux"

	handlerpackage "github.com/Rob-a21/enjoy_recipe_backend/GoLang-Backend-/api/delivery/http/handler"
	"github.com/Rob-a21/enjoy_recipe_backend/GoLang-Backend-/api/entity"
	er "github.com/Rob-a21/enjoy_recipe_backend/GoLang-Backend-/api/recipe/repository"
	es "github.com/Rob-a21/enjoy_recipe_backend/GoLang-Backend-/api/recipe/services"
	"github.com/jinzhu/gorm"

	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"

	_ "github.com/lib/pq"
)

//func createTables(dbConn *gorm.DB) []error {
//	dbConn.DropTableIfExists(&entity.Role{},&entity.User{}, &entity.Election{},
//		&entity.Party{}).GetErrors()
//	errs := dbConn.CreateTable(&entity.Role{},&entity.User{},&entity.Election{}, &entity.Party{}).GetErrors()
//	dbConn.Debug().Model(&entity.User{}).AddForeignKey("user_id", "user(Id)", "cascade", "cascade")
//	dbConn.Debug().Model(&entity.Role{}).AddForeignKey("role_id", "role(Id)", "cascade", "cascade")
//	dbConn.Debug().Model(&entity.Election{}).AddForeignKey("election_id", "election(Id)", "cascade", "cascade")
//	dbConn.Debug().Model(&entity.Party{}).AddForeignKey("party_id", "party(Id)", "cascade", "cascade")
//
//
//	if len(errs )>0 {
//		return errs
//	}
//	return nil
//}

func createTables(dbConn *gorm.DB) []error {

	dbConn.DropTableIfExists(&entity.Recipe{}).GetErrors()

	err := dbConn.CreateTable(&entity.Recipe{}).GetErrors()

	if len(err) > 0 {
		return err
	}

	return nil
}

func main() {

	//dbconn, err := gorm.Open("postgres",
	//	"postgres://postgres:postgres:admin@localhost/parties?sslmode=disable")

	dbconn, err := gorm.Open("postgres", " port=5432 user=postgres dbname=postgres sslmode=disable password=postgres")
	if dbconn != nil {
		defer dbconn.Close()
	}
	if err != nil {
		panic(err)
	}
	createTables(dbconn)

	router := mux.NewRouter()

	RecipeRepo := er.NewRecipeGormRepo(dbconn)
	RecipeService := es.NewRecipeService(RecipeRepo)
	RecipeHandler := handlerpackage.NewRecipeApiHandler(RecipeService)

	router.HandleFunc("/recipies", RecipeHandler.GetRecipies).Methods("GET")
	router.HandleFunc("/recipies/{id}", RecipeHandler.GetRecipe).Methods("GET")
	router.HandleFunc("/recipies", RecipeHandler.PostRecipe).Methods("POST")
	router.HandleFunc("/recipies/{id}", RecipeHandler.UpdateRecipe).Methods("PUT")
	router.HandleFunc("/recipies/{id}", RecipeHandler.DeleteRecipe).Methods("DELETE")

	fmt.Println("listening at PORT:9090...")

	err = http.ListenAndServe("localhost:9090", router)

	if err != nil {
		panic(err)
	}

}
