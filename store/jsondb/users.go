package jsondb

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/datnguyen-dev/tokoin_test/common"
	"github.com/datnguyen-dev/tokoin_test/store"
)

//User - Implement User Store
type user struct {
	contend        []byte
	users          []*store.User
	mapJSONField   map[string]string
	mapJSONType    map[string]string
	mapValueSearch map[string][]*store.User
}

//InitData - User
func (o *user) InitData() (bool, error) {
	if len(o.contend) > 0 {
		o.users = make([]*store.User, 0)
		err := json.Unmarshal(o.contend, &o.users)
		if err == nil {
			o.mapValueSearch = make(map[string][]*store.User)
			for idx, usr := range o.users {
				if idx < 1 {
					o.mapJSONField, o.mapJSONType = common.Mapping(usr)
				}
				for key, field := range o.mapJSONField {
					//for array
					if typeVal, ok := o.mapJSONType[key]; ok && strings.HasPrefix(typeVal, "[]") {
						vals := o.getValuesField(usr, field)
						for _, v := range vals {
							fieldCombine := fmt.Sprintf("%v[%v]", key, strings.ToLower(v))
							o.mapValueSearch[fieldCombine] = append(o.mapValueSearch[fieldCombine], usr)
						}
					} else { //for single data
						val := o.getValueField(usr, field)
						fieldCombine := fmt.Sprintf("%v[%v]", key, strings.ToLower(val))
						o.mapValueSearch[fieldCombine] = append(o.mapValueSearch[fieldCombine], usr)
					}
				}
			}
			return true, nil
		}
		return false, fmt.Errorf("Marshal organization data: " + err.Error())
	}
	return false, fmt.Errorf("Cannot load organization data")
}

//GetSearchFields -
func (o *user) GetSearchFields() ([]string, error) {
	if o.mapJSONField != nil {
		res := []string{}
		for key := range o.mapJSONField {
			res = append(res, key)
		}
		return res, nil
	}
	return nil, fmt.Errorf("not found")
}

//Search -
func (o *user) Search(field, value string) ([]*store.User, error) {
	if res, ok := o.mapValueSearch[fmt.Sprintf("%v[%v]", field, strings.ToLower(value))]; ok {
		return res, nil
	}
	return nil, fmt.Errorf("Not found")
}

func (o *user) getValueField(v *store.User, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return fmt.Sprintf("%v", f.Interface())
}

func (o *user) getValuesField(v *store.User, field string) []string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Interface().([]string)
}
