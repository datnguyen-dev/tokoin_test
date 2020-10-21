package json

import "github.com/datnguyen-dev/tokoin_test/store"

//User - Implement User Store
type User struct{}

//AddNew user json data
func (u *User) AddNew(item *store.User) (*store.User, error) {
	return nil, nil
}

//Delete - user json data
func (u *User) Delete(id int) (bool, error) {
	return false, nil
}

//Update user json data
func (u *User) Update(item *store.User) (bool, error) {
	return false, nil
}

//GetUserByID - user json data via ID
func (u *User) GetUserByID(id int) ([]*store.User, error) {
	return nil, nil
}
