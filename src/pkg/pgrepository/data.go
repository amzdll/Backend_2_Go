package pgrepository

import (
	"fmt"
	"reflect"
)

type tableData struct {
	columns      []string
	values       []interface{}
	placeholders []string
}

func checkData(src any) error {
	if reflect.TypeOf(src).Kind() != reflect.Struct {
		return fmt.Errorf("incorrect datatype")
	}
	return nil
}

func parseData(data any) tableData {
	var td tableData
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get("db")
		if tag == "" {
			continue
		}
		td.columns = append(td.columns, tag)
		td.placeholders = append(td.placeholders, fmt.Sprintf("$%d", len(td.placeholders)+1))
		td.values = append(td.values, value.Interface())
	}
	return td
}
