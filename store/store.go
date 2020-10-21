package store

import "errors"

var (
	// ErrNotFound means the requested item is not found.
	ErrNotFound = errors.New("store: item not found")
	// ErrConflict means the operation failed because of a conflict between items.
	ErrConflict = errors.New("store: item conflict")
	// ErrBeingUsed means the requested item is being used.
	ErrBeingUsed = errors.New("store: item being used")
)

//Store - Instruction
type Store interface {
	OrganizationStore()
	TicketStore()
	UserStore()
}

//OrganizationStore - Define methods for store Organization
type OrganizationStore interface {
	AddNew(item *Organization) (*Organization, error)
	Delete(id int) (bool, error)
	Update(item *Organization) (bool, error)
	GetOrganizationByID(id int) ([]*Organization, error)
}

//TicketStore - Define methods for store Ticket
type TicketStore interface {
	AddNew(item *Ticket) (*Ticket, error)
	Delete(id int) (bool, error)
	Update(item *Ticket) (bool, error)
	GetTicketByID(id int) ([]*Ticket, error)
}

//UserStore - Define methods for store User
type UserStore interface {
	AddNew(item *User) (*User, error)
	Delete(id int) (bool, error)
	Update(item *User) (bool, error)
	GetUserByID(id int) ([]*User, error)
}
