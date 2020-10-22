package jsondb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ByOrg = (`[
		{
			"_id": 101,
			"url": "http://initech.tokoin.io.com/api/v2/organizations/101.json",
			"external_id": "9270ed79-35eb-4a38-a46f-35725197ea8d",
			"name": "Enthaze",
			"domain_names": [
				"kage.com",
				"ecratic.com",
				"endipin.com",
				"zentix.com"
			],
			"created_at": "2016-05-21T11:10:28 -10:00",
			"details": "MegaCorp",
			"shared_tickets": false,
			"tags": [
				"Fulton",
				"West",
				"Rodriguez",
				"Farley"
			]
		},
		{
			"_id": 102,
			"url": "http://initech.tokoin.io.com/api/v2/organizations/102.json",
			"external_id": "7cd6b8d4-2999-4ff2-8cfd-44d05b449226",
			"name": "Nutralab",
			"domain_names": [
				"trollery.com",
				"datagen.com",
				"bluegrain.com",
				"dadabase.com"
			],
			"created_at": "2016-04-07T08:21:44 -10:00",
			"details": "Non profit",
			"shared_tickets": false,
			"tags": [
				"Cherry",
				"Collier",
				"Fuentes",
				"Trevino"
			]
		},
		{
			"_id": 103,
			"url": "http://initech.tokoin.io.com/api/v2/organizations/103.json",
			"external_id": "e73240f3-8ecf-411d-ad0d-80ca8a84053d",
			"name": "Plasmos",
			"domain_names": [
				"comvex.com",
				"automon.com",
				"verbus.com",
				"gogol.com"
			],
			"created_at": "2016-05-28T04:40:37 -10:00",
			"details": "Non profit",
			"shared_tickets": false,
			"tags": [
				"Parrish",
				"Lindsay",
				"Armstrong",
				"Vaughn"
			]
		}
	]`)
)

func TestOrgInitData(t *testing.T) {
	org := organization{contend: []byte(ByOrg)}
	bol, _ := org.InitData()
	assert.True(t, bol)
	assert.True(t, org.mapJSONField != nil)
	assert.True(t, org.mapJSONType != nil)
	assert.True(t, org.mapValueSearch != nil)
}

func TestOrgSearchField(t *testing.T) {
	org := organization{contend: []byte(ByOrg)}
	org.InitData()
	res, err := org.GetSearchFields()
	assert.True(t, len(res) > 0 && err == nil)
}

func TestOrgSearch(t *testing.T) {
	org := organization{contend: []byte(ByOrg)}
	org.InitData()
	//test with id exsited
	res, err := org.Search("_id", "103")
	assert.True(t, len(res) > 0 && err == nil)
	res, err = org.Search("_id", "3333")
	assert.True(t, err != nil)
	//test with name
	res, err = org.Search("name", "Plasmos")
	assert.True(t, len(res) > 0 && err == nil)
	res, err = org.Search("name", "abdsdf")
	assert.True(t, err != nil)
	//test with tag
	res, err = org.Search("tags", "Lindsay")
	assert.True(t, len(res) > 0 && err == nil)
	res, err = org.Search("tags", "abdsdf")
	assert.True(t, err != nil)
}
