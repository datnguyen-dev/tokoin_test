package main

import (
	"fmt"

	"github.com/datnguyen-dev/tokoin_test/cmd"
	"github.com/datnguyen-dev/tokoin_test/config"
	"github.com/datnguyen-dev/tokoin_test/store/jsondb"
)

func main() {
	AppName := "tokoin_test"
	ConfigFile := fmt.Sprintf("%s.conf", AppName)
	cnf, err := config.InitConfig(ConfigFile)
	if err != nil {
		panic(err)
	}
	fmt.Println(cnf.JSONDb.OrgPath)
	store := jsondb.Connect(cnf.JSONDb.OrgPath, cnf.JSONDb.TicketPath, cnf.JSONDb.UserPath)
	cmd.Execute(store)
}
