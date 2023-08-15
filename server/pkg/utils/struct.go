package utils

import (
	"github.com/iancoleman/strcase"
	"reflect"
)

func FillStruct(data map[string]interface{}, result interface{}) {
	t := reflect.ValueOf(result).Elem()
	for k, v := range data {
		val := t.FieldByName(strcase.ToCamel(k))
		val.Set(reflect.ValueOf(v))
	}
}
