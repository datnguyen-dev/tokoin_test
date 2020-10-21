package json

import "github.com/datnguyen-dev/tokoin_test/store"

//Organization - Implement Organization Store
type Organization struct{}

//AddNew - AddNew Organization Json store
func (o *Organization) AddNew(item *store.Organization) (*store.Organization, error) {
	return nil, nil
}

//Delete - Delete Organization Json Data
func (o *Organization) Delete(id int) (bool, error) {
	return false, nil
}

//Update - Update Organization Json Data
func (o *Organization) Update(item *store.Organization) (bool, error) {
	return false, nil
}

//GetOrganizationByID - Get Organization Json Data by id
func (o *Organization) GetOrganizationByID(id int) ([]*store.Organization, error) {
	return nil, nil
}
