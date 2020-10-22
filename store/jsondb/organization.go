package jsondb

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/datnguyen-dev/tokoin_test/common"
	"github.com/datnguyen-dev/tokoin_test/store"
)

//Organization - Implement Organization Store
type organization struct {
	contend        []byte
	organizations  []*store.Organization
	mapJSONField   map[string]string
	mapJSONType    map[string]string
	mapValueSearch map[string][]*store.Organization
}

//InitData - Organization
func (o *organization) InitData() (bool, error) {
	if len(o.contend) > 0 {
		o.organizations = make([]*store.Organization, 0)
		err := json.Unmarshal(o.contend, &o.organizations)
		if err == nil {
			o.mapValueSearch = make(map[string][]*store.Organization)
			for idx, org := range o.organizations {
				if idx < 1 {
					o.mapJSONField, o.mapJSONType = common.Mapping(org)
				}
				for key, field := range o.mapJSONField {
					//for array
					if typeVal, ok := o.mapJSONType[key]; ok && strings.HasPrefix(typeVal, "[]") {
						vals := o.getValuesField(org, field)
						for _, v := range vals {
							fieldCombine := fmt.Sprintf("%v[%v]", key, strings.ToLower(v))
							o.mapValueSearch[fieldCombine] = append(o.mapValueSearch[fieldCombine], org)
						}
					} else { //for single data
						val := o.getValueField(org, field)
						fieldCombine := fmt.Sprintf("%v[%v]", key, strings.ToLower(val))
						o.mapValueSearch[fieldCombine] = append(o.mapValueSearch[fieldCombine], org)
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
func (o *organization) GetSearchFields() ([]string, error) {
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
func (o *organization) Search(field, value string) ([]*store.Organization, error) {
	if res, ok := o.mapValueSearch[fmt.Sprintf("%v[%v]", field, strings.ToLower(value))]; ok {
		return res, nil
	}
	return nil, fmt.Errorf("Not found")
}

func (o *organization) getValueField(v *store.Organization, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return fmt.Sprintf("%v", f.Interface())
}

func (o *organization) getValuesField(v *store.Organization, field string) []string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Interface().([]string)
}
