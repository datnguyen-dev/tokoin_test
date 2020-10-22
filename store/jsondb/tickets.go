package jsondb

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/datnguyen-dev/tokoin_test/common"
	"github.com/datnguyen-dev/tokoin_test/store"
)

//Ticket - Implement Ticket Store
type ticket struct {
	contend        []byte
	tickets        []*store.Ticket
	mapJSONField   map[string]string
	mapJSONType    map[string]string
	mapValueSearch map[string][]*store.Ticket
}

//InitData - Ticket
func (o *ticket) InitData() (bool, error) {
	if len(o.contend) > 0 {
		o.tickets = make([]*store.Ticket, 0)
		err := json.Unmarshal(o.contend, &o.tickets)
		if err == nil {
			o.mapValueSearch = make(map[string][]*store.Ticket)
			for idx, tick := range o.tickets {
				if idx < 1 {
					o.mapJSONField, o.mapJSONType = common.Mapping(tick)
				}
				for key, field := range o.mapJSONField {
					//for array
					if typeVal, ok := o.mapJSONType[key]; ok && strings.HasPrefix(typeVal, "[]") {
						vals := o.getValuesField(tick, field)
						for _, v := range vals {
							fieldCombine := fmt.Sprintf("%v[%v]", key, strings.ToLower(v))
							o.mapValueSearch[fieldCombine] = append(o.mapValueSearch[fieldCombine], tick)
						}
					} else { //for single data
						val := o.getValueField(tick, field)
						fieldCombine := fmt.Sprintf("%v[%v]", key, strings.ToLower(val))
						o.mapValueSearch[fieldCombine] = append(o.mapValueSearch[fieldCombine], tick)
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
func (o *ticket) GetSearchFields() ([]string, error) {
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
func (o *ticket) Search(field, value string) ([]*store.Ticket, error) {
	if res, ok := o.mapValueSearch[fmt.Sprintf("%v[%v]", field, strings.ToLower(value))]; ok {
		return res, nil
	}
	return nil, fmt.Errorf("Not found")
}

func (o *ticket) getValueField(v *store.Ticket, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return fmt.Sprintf("%v", f.Interface())
}

func (o *ticket) getValuesField(v *store.Ticket, field string) []string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Interface().([]string)
}
