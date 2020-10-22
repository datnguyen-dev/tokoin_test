package main

import (
	"github.com/datnguyen-dev/tokoin_test/cmd"
	"github.com/datnguyen-dev/tokoin_test/store/jsondb"
)

func main() {
	pathOrg := "./resources/organizations.json"
	pathTicket := "./resources/tickets.json"
	pathUser := "./resources/users.json"

	store := jsondb.Connect(pathOrg, pathTicket, pathUser)
	cmd.Execute(store)
}
