package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapping(t *testing.T) {
	type Organization struct {
		ID            int      `json:"_id"`
		URL           string   `json:"url"`
		ExternalID    string   `json:"external_id"`
		Name          string   `json:"name"`
		DomainNames   []string `json:"domain_names"`
		CreateAt      string   `json:"created_at"`
		Details       string   `json:"details"`
		SharedTickets bool     `json:"shared_tickets"`
		Tags          []string `json:"tags"`
	}
	orgs := new(Organization)
	orgs.ID = 1
	orgs.Details = "1111"
	orgs.URL = "https://www.vietnames.vn"

	mapField := make(map[string]string)
	mapType := make(map[string]string)
	mapField, mapType = Mapping(orgs)
	assert.True(t, mapField != nil)
	assert.True(t, mapType != nil)
}
