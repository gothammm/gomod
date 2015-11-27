package gomod

import (
	"errors"
	"reflect"
	"strconv"
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
			val := field.String()

			if _, ok := rules[RequiredTag]; !ok && field.Interface() == reflect.Zero(field.Type()).Interface() {
				continue
			}

			if j != nil {

				switch i {
				case EmailTag:
					if !emailRX.MatchString(val) {
						errorsList = append(errorsList, &ModError{Message: "Invalid Email", FieldName: k.Name, ModelName: m.Type().Name()})
					}
				case RequiredTag:
					if field.Interface() == reflect.Zero(field.Type()).Interface() {
						errorsList = append(errorsList, &ModError{Message: "This field is required", FieldName: k.Name, ModelName: m.Type().Name()})
					}
				case PhoneIndiaTag:
					if !INPhoneRX.MatchString(val) {
						errorsList = append(errorsList, &ModError{Message: "Invalid Phone Number", FieldName: k.Name, ModelName: m.Type().Name()})
					}

				case MaxLenTag:
					maxLen := j.(int)
					if field.Kind() == reflect.Int && maxLen < field.Interface().(int) {
						errorsList = append(errorsList, &ModError{Message: "length cannot exceed " + strconv.Itoa(maxLen), FieldName: k.Name, ModelName: m.Type().Name()})
					} else if field.Kind() != reflect.Int && maxLen < field.Len() {
						errorsList = append(errorsList, &ModError{Message: "length cannot exceed " + strconv.Itoa(maxLen), FieldName: k.Name, ModelName: m.Type().Name()})

					}

				case MinLenTag:
					minLen := j.(int)
					if field.Kind() == reflect.Int && minLen > field.Interface().(int) {
						errorsList = append(errorsList, &ModError{Message: "length cannot be less than " + strconv.Itoa(minLen), FieldName: k.Name, ModelName: m.Type().Name()})
					} else if field.Kind() != reflect.Int && minLen > field.Len() {
						errorsList = append(errorsList, &ModError{Message: "length cannot be less than " + strconv.Itoa(minLen), FieldName: k.Name, ModelName: m.Type().Name()})
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
