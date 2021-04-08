package services

import (
	party "github.com/Rob-a21/flutter_backend/api/party"
	"github.com/Rob-a21/flutter_backend/api/entity"
)

type PartyService struct {
	partyRepo party.PartyRepository
}

func NewPartyService(partyRepo party.PartyRepository) *PartyService {
	return &PartyService{partyRepo: partyRepo}
}

func (cs *PartyService) Parties() ([]entity.Party, []error) {

	parties, errs := cs.partyRepo.Parties()
	if len(errs) > 0 {
		return nil, errs
	}
	return parties, errs

}

func (cs *PartyService) Party(id uint32) (*entity.Party, []error) {
	par, errs := cs.partyRepo.Party(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return par, errs

}

func (cs *PartyService) UpdateParty(party *entity.Party) (*entity.Party, []error) {
	par, errs := cs.partyRepo.UpdateParty(party)
	if len(errs) > 0 {
		return nil, errs
	}
	return par, errs

}

func (cs *PartyService) DeleteParty(id uint32) (*entity.Party, []error) {

	par, errs := cs.partyRepo.DeleteParty(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return par, errs
}

func (cs *PartyService) StoreParty(party *entity.Party) (*entity.Party, []error) {

	par, errs := cs.partyRepo.StoreParty(party)
	if len(errs) > 0 {
		return nil, errs
	}
	return par, errs
}
