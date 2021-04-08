package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/Rob-a21/flutter_backend/api/entity"
)

type PartyGormRepo struct {
	conn *gorm.DB
}

func NewPartyGormRepo(db *gorm.DB) *PartyGormRepo {
	return &PartyGormRepo{conn: db}
}

func (partyRepo *PartyGormRepo) Parties() ([]entity.Party, []error) {
	var parties []entity.Party
	errs := partyRepo.conn.Find(&parties).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return parties, errs
}

func (partyRepo *PartyGormRepo) Party(id uint32) (*entity.Party, []error) {
	party := entity.Party{}
	errs := partyRepo.conn.First(&party, id).GetErrors()
	if len(errs) > 0 {
		//fmt.Println(errs)
		return nil, errs
	}
	return &party, errs
}

//func (partyRepo *PartyGormRepo) UpdateParty(party *entity.Party) (*entity.Party, []error) {
//	errs := partyRepo.conn.Save(party).GetErrors()
//	if len(errs) > 0 {
//		return nil, errs
//	}
//	return party, errs
//}
   // =========================================================
func (partyRepo *PartyGormRepo) UpdateParty(party *entity.Party) (*entity.Party, []error) {
	errs := partyRepo.conn.Save(party).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return party, errs
}
func (partyRepo *PartyGormRepo) DeleteParty(id uint32) (*entity.Party, []error) {
	party, errs := partyRepo.Party(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = partyRepo.conn.Delete(party, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return party, errs
}

func (partyRepo *PartyGormRepo) StoreParty(party *entity.Party) (*entity.Party, []error) {
	errs := partyRepo.conn.Create(party).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return party, errs
}
