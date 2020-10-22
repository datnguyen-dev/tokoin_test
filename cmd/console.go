package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/datnguyen-dev/tokoin_test/common"
	"github.com/datnguyen-dev/tokoin_test/store"
)

//Execute - run cmd
func Execute(store store.Store) {
	loadData(store)
	cont := true
	for cont {
		introduction()
		sel, err := common.StringInput("Please select")
		if err != nil {
			fmt.Println("Input error: " + err.Error())
			continue
		}
		switch sel {
		case "quit":
		case "q":
			cont = false
			fmt.Println("bye!!!!")
			break
		case "1":
			search(store)
			break
		case "2":
			searchableFields(store)
			break
		default:
			fmt.Println("Selected option is invalid. Please try again.")
		}
	}
}

func search(store store.Store) {
	fmt.Println("Select 1) Users or 2) Tickets or 3) Organizations")
	sel, err := common.StringInput("Please select")
	if err != nil {
		fmt.Println("Input error")
		return
	}
	switch sel {
	case "1":
		searchUser(store)
		break
	case "2":
		searchTicket(store)
		break
	case "3":
		searchOrganization(store)
		break
	default:
		fmt.Println("Selected option is invalid.")
	}
}

func searchableFields(store store.Store) {
	fmt.Println("Select 1) Users or 2) Tickets or 3) Organizations")
	sel, err := common.StringInput("Please select")
	if err != nil {
		fmt.Println("Input error")
		return
	}
	switch sel {
	case "1":
		vals, err := store.Users().GetSearchFields()
		if err != nil {
			fmt.Println("No data found: " + err.Error())
			return
		}
		fmt.Println("Search field for user:")
		printArray(vals)
		return
	case "2":
		vals, err := store.Tickets().GetSearchFields()
		if err != nil {
			fmt.Println("No data found")
			return
		}
		fmt.Println("Search field for tickets:")
		printArray(vals)
		return
	case "3":
		vals, err := store.Organizations().GetSearchFields()
		if err != nil {
			fmt.Println("No data found")
			return
		}
		fmt.Println("Search field for organizations:")
		printArray(vals)
		return
	default:
		fmt.Println("Selected option is invalid.")
	}
}

func introduction() {
	fmt.Println()
	fmt.Println("Type 'quit' or 'q' to exit at any time, Press 'Enter' to continue")
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("Select search options:")
	fmt.Println("** Press 1 to search")
	fmt.Println("** Press 2 to view a list of searchable fields")
	fmt.Println("** Type quit to exit")
	fmt.Println("-----------------------------------------------------------------")
}

func printArray(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func loadData(store store.Store) {
	//load organization
	res, err := store.Organizations().InitData()
	if !res {
		panic(err)
	}
	//load ticket
	res, err = store.Tickets().InitData()
	if !res {
		panic(err)
	}
	//load user
	res, err = store.Users().InitData()
	if !res {
		panic(err)
	}
}

func searchUser(store store.Store) {
	//Searching users MUST return his/her assignee ticket subject and submitted ticket subject and his/her organization name
	key, val, err := inputKeyVal()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("-------------------Results Users-------------------")
	usrs, err := store.Users().Search(key, val)
	if err != nil {
		fmt.Println(err)
		goto End
	}
	for _, usr := range usrs {
		//print user first
		s, _ := json.MarshalIndent(usr, "", " ")
		fmt.Println(string(s))
		//add ticket assignee
		tiks, err := store.Tickets().Search("assignee_id", strconv.Itoa(usr.ID))
		if err == nil {
			for idx, tik := range tiks {
				fmt.Println(fmt.Sprintf(`Assignee_Tiket_%v: "subject": %s`, idx, tik.Subject))
			}
		} else {
			fmt.Println(fmt.Sprintf(`Assignee_Tiket_%v:`, 0))
		}
		//add ticket submited
		tiks, err = store.Tickets().Search("submitter_id", strconv.Itoa(usr.ID))
		if err == nil {
			for idx, tik := range tiks {
				fmt.Println(fmt.Sprintf(`Submited_Tiket_%v: "subject": %s`, idx, tik.Subject))
			}
		} else {
			fmt.Println(fmt.Sprintf(`Submited_Tiket_%v:`, 0))
		}
	}
End:
	fmt.Println("-------------------End-------------------")
	fmt.Print("Press any key to continue")
	common.StringInput("")
	common.ClearInputScreen()
}

func searchTicket(store store.Store) {
	//Searching tickets MUST return its assignee name, submitter name, and organization name.
	key, val, err := inputKeyVal()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("-------------------Results Ticket-------------------")
	ticks, err := store.Tickets().Search(key, val)
	if err != nil {
		fmt.Println(err)
		goto End
	}
	for _, tick := range ticks {
		//print ticket first
		s, _ := json.MarshalIndent(tick, "", " ")
		fmt.Println(string(s))
		//add assignee name
		users, err := store.Users().Search("_id", strconv.Itoa(tick.AssigneeID))
		if err == nil {
			for _, usr := range users {
				fmt.Println(fmt.Sprintf(`Assignee: "name": %s`, usr.Name))
			}
		} else {
			fmt.Println(`Assignee:`)
		}
		//add submit name
		users, err = store.Users().Search("_id", strconv.Itoa(tick.SubmitterID))
		if err == nil {
			for _, usr := range users {
				fmt.Println(fmt.Sprintf(`Submitter: "name": %s`, usr.Name))
			}
		} else {
			fmt.Println(`Submitter:`)
		}
		//add organization name
		orgs, err := store.Organizations().Search("organization_id", strconv.Itoa(tick.OrganizationID))
		if err == nil {
			for _, org := range orgs {
				fmt.Println(fmt.Sprintf(`Organization: "name": %s`, org.Name))
			}
		} else {
			fmt.Println(`Organization:`)
		}
	}
End:
	fmt.Println("-------------------End-------------------")
	fmt.Print("Press any key to continue")
	common.StringInput("")
	common.ClearInputScreen()
}

func searchOrganization(store store.Store) {
	//Searching organization MUST return its ticket subject and users name
	key, val, err := inputKeyVal()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("-------------------Results Organization-------------------")
	orgs, err := store.Organizations().Search(key, val)
	if err != nil {
		fmt.Println(err)
		goto End
	}
	for _, org := range orgs {
		//print organization first
		s, _ := json.MarshalIndent(org, "", " ")
		fmt.Println(string(s))
		//add ticket subject
		tiks, err := store.Tickets().Search("organization_id", strconv.Itoa(org.ID))
		if err == nil {
			for idx, tik := range tiks {
				fmt.Println(fmt.Sprintf(`Tiket_%v: "subject": %s`, idx, tik.Subject))
			}
		} else {
			fmt.Println(fmt.Sprintf(`Tiket_%v:`, 0))
		}
		//add user name
		usrs, err := store.Users().Search("organization_id", strconv.Itoa(org.ID))
		if err == nil {
			for idx, usr := range usrs {
				fmt.Println(fmt.Sprintf(`User_%v: "name": %s`, idx, usr.Name))
			}
		} else {
			fmt.Println(fmt.Sprintf(`User_%v:`, 0))
		}
	}
End:
	fmt.Println("-------------------End-------------------")
	fmt.Print("Press any key to continue")
	common.StringInput("")
	common.ClearInputScreen()
}

func inputKeyVal() (string, string, error) {
	key, err := common.StringInput("Input key: ")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	val, err := common.StringInput("Input Value: ")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	return key, val, nil
}
