package gomod

import (
	"reflect"
)

type Field struct {
	Name            string
	JsonName        string
	ValidationRules map[string]interface{}
}

func getFields(m reflect.Value) ([]*Field, error) {
	var fields []*Field

	if m.Kind() == reflect.Ptr {
		m = m.Elem()
	}

	t := m.Type()

	fieldCount := t.NumField()

	for i := 0; i < fieldCount; i++ {
		field := t.Field(i)

		validationRule, err := getValidationRules(field)

		if err != nil {
			return nil, err
		}
		fields = append(fields, &Field{Name: field.Name, JsonName: getFieldName(field), ValidationRules: validationRule})
	}
	return fields, nil
}

func Fields(m reflect.Value) ([]*Field, error) {
	return getFields(m)
}
