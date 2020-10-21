package json

import "github.com/datnguyen-dev/tokoin_test/store"

//Ticket - Implement Ticket Store
type Ticket struct{}

//AddNew ticket json data
func (t *Ticket) AddNew(item *store.Ticket) (*store.Ticket, error) {
	return nil, nil
}

//Delete - ticket json data
func (t *Ticket) Delete(id int) (bool, error) {
	return false, nil
}

//Update ticket json data
func (t *Ticket) Update(item *store.Ticket) (bool, error) {
	return false, nil
}

//GetTicketByID - ticket json data via ID
func (t *Ticket) GetTicketByID(id int) ([]*store.Ticket, error) {
	return nil, nil
}
