package main

import (
	"fmt"

	"github.com/gorilla/mux"

	"github.com/Rob-a21/flutter_backend/api/delivery/http/handler"
	er "github.com/Rob-a21/flutter_backend/api/election/repository"
	es "github.com/Rob-a21/flutter_backend/api/election/service"
	"github.com/Rob-a21/flutter_backend/api/entity"
	cr "github.com/Rob-a21/flutter_backend/api/party/repository"
	cs "github.com/Rob-a21/flutter_backend/api/party/services"
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
	//dbConn.DropTableIfExists(&entity.Role{},&entity.User{}).GetErrors()
	//errs := dbConn.CreateTable(&entity.Role{},&entity.User{}).GetErrors()
	//dbConn.Debug().Model(&entity.User{}).AddForeignKey("user_id", "user(Id)", "cascade", "cascade")
	//dbConn.Debug().Model(&entity.Role{}).AddForeignKey("role_id", "role(Id)", "cascade", "cascade")

	//dbConn.DropTableIfExists(&entity.User{}).GetErrors()
	//dbConn.DropTableIfExists(&entity.Role{}).GetErrors()
	//	dbConn.Debug().Model(&entity.User{}).AddForeignKey("user_id", "user(Id)", "cascade", "cascade")
	//	dbConn.Debug().Model(&entity.Role{}).AddForeignKey("role_id", "role(Id)", "cascade", "cascade")
	dbConn.DropTableIfExists(&entity.Party{}).GetErrors()
	dbConn.DropTableIfExists(&entity.Election{}).GetErrors()
	//
	//rl := dbConn.CreateTable(&entity.Role{}).GetErrors()
	//us := dbConn.CreateTable(&entity.User{}).GetErrors()
	err := dbConn.CreateTable(&entity.Party{}).GetErrors()
	er := dbConn.CreateTable(&entity.Election{}).GetErrors()

	 if len(er) >0{
		return er
	} else if len(err) >0{
		return err}

	return nil
}


func main() {

	//dbconn, err := gorm.Open("postgres",
	//	"postgres://postgres:postgres:admin@localhost/parties?sslmode=disable")

	dbconn, err := gorm.Open( "postgres", " port=5432 user=postgres dbname=postgres sslmode=disable password=postgres")
	if dbconn != nil {
		defer dbconn.Close()
	}
	if err != nil {
		panic(err)
	}
	createTables(dbconn)

	router := mux.NewRouter()

	PartyRepo := cr.NewPartyGormRepo(dbconn)
	PartyService := cs.NewPartyService(PartyRepo)
	partyHandler := handlerpackage.NewPartyApiHandler(PartyService)

	ElectionRepo := er.NewElectionGormRepo(dbconn)
	ElectionService := es.NewElectionService(ElectionRepo)
	electionHandler := handlerpackage.NewElectionApiHandler(ElectionService)

	//roleGormRepo :=kr.NewRoleGormRepo(dbconn)
	//roleService :=ks.NewRoleService(roleGormRepo)
	//roleHandler:=handlerpackage.NewRoleApiHandler(roleService)
	//
	//usersRepo := kr.NewUserGormRepo(dbconn)
	//usersService := ks.NewUserService(usersRepo)
	//usersHandler := handlerpackage.NewUserApiHandler(usersService)

	//router.HandleFunc("/v1/role", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.GetRoles))).Methods("GET")
	//router.HandleFunc("/v1/roles/{name}", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.GetRoleByName))).Methods("GET")
	//router.HandleFunc("/v1/role/{id}",usersHandler.Authenticated(usersHandler.Authorized(roleHandler.GetRoleByID))).Methods("GET")
	//router.HandleFunc("/v1/role",roleHandler.PostRole).Methods("POST")
	//router.HandleFunc("/v1/role/{id}", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.PutRole))).Methods("PUT")
	//router.HandleFunc("/v1/role/{id}", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.DeleteRole))).Methods("DELETE")
	//
	//
	//router.HandleFunc("/v1/admin/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUser))).Methods("GET")
	//router.HandleFunc("/v1/admin/users", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUsers))).Methods("GET")
	//router.HandleFunc("/v1/admin/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.PutUser))).Methods("PUT")
	//router.HandleFunc("/v1/admin/users", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.PostUser))).Methods("POST")
	//router.HandleFunc("/v1/admin/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.DeleteUser))).Methods("DELETE")
	//router.HandleFunc("/v1/admin/email/{email}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.IsEmailExists))).Methods("GET")
	//router.HandleFunc("/v1/admin/phone/{phone}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.IsPhoneExists))).Methods("GET")
	//router.HandleFunc("/v1/admin/check", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUserByUsernameAndPassword))).Methods("POST")
	//
	//router.HandleFunc("/v1/user/users", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.PostUser))).Methods("POST")
	//router.HandleFunc("/v1/user/users", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUsers))).Methods("GET")
	//router.HandleFunc("/v1/user/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUser))).Methods("GET")
	//router.HandleFunc("/v1/user/users/{id}", usersHandler.Authenticated(usersHandler.PutUser)).Methods("PUT")
	//router.HandleFunc("/v1/user/password/{id}", usersHandler.Authenticated(usersHandler.ChangePassword)).Methods("PUT")
	//router.HandleFunc("/v1/user/password/{id}", usersHandler.Authenticated(usersHandler.CheckPassword)).Methods("POST")
	//router.HandleFunc("/v1/user/email/{email}", usersHandler.IsEmailExists).Methods("GET")
	//router.HandleFunc("/v1/user/phone/{phone}", usersHandler.IsPhoneExists).Methods("GET")
	//router.HandleFunc("/v1/user/check", usersHandler.GetUserByUsernameAndPassword).Methods("POST")
	//router.HandleFunc("/v1/user/login", usersHandler.Login).Methods("POST")
	//router.HandleFunc("/v1/user/signup", usersHandler.SignUp).Methods("POST")

	router.HandleFunc("/parties", partyHandler.GetParties).Methods("GET")
	router.HandleFunc("/parties/{id}", partyHandler.GetParty).Methods("GET")
	router.HandleFunc("/parties", partyHandler.PostParty).Methods("POST")
	router.HandleFunc("/parties/{id}", partyHandler.UpdateParty).Methods("PUT")
	router.HandleFunc("/parties/{id}", partyHandler.DeleteParty).Methods("DELETE")

	router.HandleFunc("/elections", electionHandler.GetElections).Methods("GET")
	router.HandleFunc("/elections/{id}", electionHandler.GetElection).Methods("GET")
	router.HandleFunc("/elections", electionHandler.PostElection).Methods("POST")
	router.HandleFunc("/elections/{id}", electionHandler.UpdateElection).Methods("PUT")
	router.HandleFunc("/elections/{id}", electionHandler.DeleteElection).Methods("DELETE")



	fmt.Println("listening at PORT:9090...")

	err = http.ListenAndServe("192.168.123.57:9090", router)

	if err != nil {
		panic(err)
	}

}
