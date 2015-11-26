package gomod

import (
	"errors"
	"reflect"
)

func validateFields(m reflect.Value, fields []*Field) []*ModError {
	if m.Kind() == reflect.Ptr {
		m = m.Elem()
	}

	var errorsList []*ModError

	for _, k := range fields {
		field := m.FieldByName(k.Name)

		rules := k.ValidationRules

		for i, j := range rules {
			if j {

				switch i {
				case EmailTag:
					val := field.String()

					if !emailRX.MatchString(val) {
						errorsList = append(errorsList, &ModError{Message: "Invalid Email", FieldName: k.Name, ModelName: m.Type().Name()})
					}
				case RequiredTag:
					if field.Interface() == reflect.Zero(field.Type()).Interface() {
						errorsList = append(errorsList, &ModError{Message: "This field is required", FieldName: k.Name, ModelName: m.Type().Name()})
					}

				}
			}
		}
	}
	return errorsList
}

func Validate(m interface{}) ([]*ModError, error) {
	if !IsStruct(m) {
		return nil, errors.New("Cannot pass non-struct to Validate")
	}
	s := reflect.ValueOf(m)

	fields, err := getFields(s)

	if err != nil {
		return nil, err
	}

	return validateFields(s, fields), nil
}

func IsStruct(m interface{}) bool {
	if m == nil {
		return false
	}
	t := reflect.TypeOf(m)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t.Kind() == reflect.Struct
}
