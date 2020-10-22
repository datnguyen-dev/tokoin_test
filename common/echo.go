package common

import (
	"fmt"
	"strings"
)

//EchoStruct - print all key of map
func EchoStruct(obj interface{}) {
	if obj == nil {
		fmt.Println("Error occure!")
	}
	var res []string
	mapobj := obj.(map[string]interface{})
	for key := range mapobj {
		res = append(res, key)
	}

	fmt.Println(strings.Join(res, ";"))
}
