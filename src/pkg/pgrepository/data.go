package pgrepository

import (
	"fmt"
	"reflect"
	"strings"
)

type Query struct {
	template     string
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

func parseCreateData(data any) Query {
	var td Query
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

func GenerateWhereQuery(filters []string, fields []string, data interface{}) (Query, error) {
	q := Query{template: "update %s "}
	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)
	var builder strings.Builder
	var values []interface{}

	builder.WriteString(" where (")
	for i, filter := range filters {
		field, found := dataType.FieldByName(filter)
		if !found {
			return q, fmt.Errorf("field %s not found in struct", filter)
		}
		fieldValue := dataValue.FieldByName(filter)
		q.values = append(q.values, fieldValue.Interface())
		builder.WriteString(fmt.Sprintf(field.Tag.Get("db")+" = $%d AND ", i+1))
	}
	q.template = q.template + builder.String()[:len(builder.String())-5]
	fmt.Println(values)
	fmt.Println(q.template)
	return q, nil
}
