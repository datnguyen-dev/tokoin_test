package jsondb

import (
	"github.com/datnguyen-dev/tokoin_test/common"
	"github.com/datnguyen-dev/tokoin_test/store"
)

//Store - define implement Store
type Store struct {
	organization *organization
	ticket       *ticket
	user         *user
}

//Organizations - implement store
func (s *Store) Organizations() store.OrganizationStore {
	return s.organization
}

//Tickets - implement Store
func (s *Store) Tickets() store.TicketStore {
	return s.ticket
}

//Users - implement store
func (s *Store) Users() store.UserStore {
	return s.user
}

//Connect -
func Connect(pathOrg, pathTicket, pathUser string) *Store {
	orgContent, err := common.ReadFileContent(pathOrg)
	if err != nil {
		panic("Load organization fail")
	}
	tickContent, err := common.ReadFileContent(pathTicket)
	if err != nil {
		panic("Load ticket fail")
	}
	usrContent, err := common.ReadFileContent(pathUser)
	if err != nil {
		panic("Load user fail")
	}
	s := &Store{
		organization: &organization{contend: orgContent},
		ticket:       &ticket{contend: tickContent},
		user:         &user{contend: usrContent},
	}
	return s
}
