package common

import (
	"fmt"
	"reflect"
)

//Mapping -
func Mapping(obj interface{}) (map[string]string, map[string]string) {
	v := reflect.ValueOf(obj).Elem()
	t := v.Type()
	mapField := make(map[string]string)
	mapType := make(map[string]string)
	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name
		fieldJSON := t.Field(i).Tag.Get("json")
		// fieldValue := v.Field(i).Interface()
		if _, ok := mapField[fieldJSON]; !ok {
			mapField[fieldJSON] = fieldName
		}
		if _, ok := mapType[fieldJSON]; !ok {
			mapType[fieldJSON] = fmt.Sprintf("%v", v.Type().Field(i).Type)
		}
		// mapValKey := fmt.Sprintf("%v[%v]", fieldJSON, strings.ToLower(fmt.Sprintf("%v", v.Field(i).Interface())))
		// if items, ok := mapValue[mapValKey]; !ok {
		// 	mapValue[mapValKey] = append(mapValue[mapValKey], obj)
		// }
	}
	return mapField, mapType
}
