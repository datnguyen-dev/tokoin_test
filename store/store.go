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
	Organizations() OrganizationStore
	Tickets() TicketStore
	Users() UserStore
}

//OrganizationStore - Define methods for store Organization
type OrganizationStore interface {
	InitData() (bool, error)
	GetSearchFields() ([]string, error)
	Search(field, value string) ([]*Organization, error)
}

//TicketStore - Define methods for store Ticket
type TicketStore interface {
	InitData() (bool, error)
	GetSearchFields() ([]string, error)
	Search(field, value string) ([]*Ticket, error)
}

//UserStore - Define methods for store User
type UserStore interface {
	InitData() (bool, error)
	GetSearchFields() ([]string, error)
	Search(field, value string) ([]*User, error)
}
