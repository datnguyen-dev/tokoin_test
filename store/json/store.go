package json

import (
	"github.com/datnguyen-dev/tokoin_test/store"
)

//Store - define implement Store
type Store struct {
	organization *Organization
	ticket       *Ticket
	user         *User
}

//OrganizationImp - implement store
func (s *Store) OrganizationImp() store.OrganizationStore {
	return s.organization
}

//TicketImp - implement Store
func (s *Store) TicketImp() store.TicketStore {
	return s.ticket
}

//UserImp - implement store
func (s *Store) UserImp() store.UserStore {
	return s.user
}
