package party

import "github.com/Rob-a21/flutter_backend/api/entity"

type PartyRepository interface {
	Parties() ([]entity.Party, []error)
	Party(id uint32) (*entity.Party, []error)
	StoreParty(user *entity.Party) (*entity.Party, []error)
	UpdateParty(order *entity.Party) (*entity.Party, []error)
	DeleteParty(id uint32) (*entity.Party, []error)
}
