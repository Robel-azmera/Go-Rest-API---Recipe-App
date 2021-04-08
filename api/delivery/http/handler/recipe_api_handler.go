package handlerpackage

import (
	"encoding/json"
	"fmt"
	"github.com/Rob-a21/flutter_backend/api/election"

	"github.com/Rob-a21/flutter_backend/api/entity"
	"github.com/Rob-a21/flutter_backend/api/utils"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ElectionApiHandler struct {
	electionServices election.ElectionServices
}

func NewElectionApiHandler(electionServices election.ElectionServices) *ElectionApiHandler {
	return &ElectionApiHandler{electionServices: electionServices}
}
func (cah *ElectionApiHandler) GetElection(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	cl, errs := cah.electionServices.Election(uint32(id))
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

func (cah *ElectionApiHandler) GetElections(w http.ResponseWriter, r *http.Request) {

	elections, errs := cah.electionServices.Elections()

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

func (cah *ElectionApiHandler) PostElection(w http.ResponseWriter, r *http.Request) {

	body := utils.BodyParser(r)
	var par entity.Election
	err := json.Unmarshal(body, &par)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	storeElection, errs := cah.electionServices.StoreElection(&par)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return

	}
	output, err := json.MarshalIndent(storeElection, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}

func (fah *ElectionApiHandler) UpdateElection(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("125")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	elct, errs := fah.electionServices.Election(uint32(id))
	var electionWithStringTimeStamp ElectionWithStringTimeStamp
	if len(errs) > 0 {
		fmt.Println("134")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	body := utils.BodyParser(r)
	err = json.Unmarshal(body, &electionWithStringTimeStamp)
	if err != nil {
		fmt.Println("153")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	elct.ID = electionWithStringTimeStamp.ID
	elct.BoardLeader = electionWithStringTimeStamp.BoardLeader
	elct.Country = electionWithStringTimeStamp.Country
	elct.Description = electionWithStringTimeStamp.Description
	elct.ElectionYear = electionWithStringTimeStamp.ElectionYear


	elct, errs = fah.electionServices.UpdateElection(elct)
	if len(errs) > 0 {
		fmt.Println("169")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(elct, "", "\t\t")

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

//func (cah *ElectionApiHandler) UpdateElection(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	id, err := strconv.Atoi(params["id"])
//	if err != nil {
//		w.Header().Set("Content-Type", "application/json")
//		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
//		return
//	}
//
//	par, errs := cah.electionServices.Election(uint32(id))
//
//	if len(errs) > 0 {
//		fmt.Println("119  %s",errs)////////////////////////////////////////////////////////////////////////////
//		w.Header().Set("Content-Type", "application/json")
//		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//		return
//	}
//	l := r.ContentLength
//
//	body := make([]byte, l)
//
//	_, err = r.Body.Read(body)
//	if err != nil {
//		//fmt.Println("130  %s",err)
//		w.Header().Set("Content-Type", "application/json")
//		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//		return
//	}
//	err = json.Unmarshal(body, &par)
//	if err != nil {
//		w.Header().Set("Content-Type", "application/json")
//		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//		return
//	}
//	par, errs = cah.electionServices.UpdateElection(par)
//
//	if len(errs) > 0 {
//		w.Header().Set("Content-Type", "application/json")
//		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//		return
//	}
//	output, err := json.MarshalIndent(par, "", "\t\t")
//
//	if err != nil {
//		w.Header().Set("Content-Type", "application/json")
//		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//	_, _ = w.Write(output)
//	return
//}
func (cah *ElectionApiHandler) DeleteElection(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := cah.electionServices.DeleteElection(uint32(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}

type ElectionWithStringTimeStamp struct {
	ID          int   `gorm:"primary_key;auto_increment" json:"id"`
	BoardLeader       string `gorm:"type:varchar(255);not null" json:"leader"`
	ElectionYear       string `gorm:"type:varchar(255);not null" json:"year"`
	Country       string `gorm:"type:varchar(255);not null" json:"country"`
	Description string `gorm:"type:varchar(255);not null" json:"description"`
}

