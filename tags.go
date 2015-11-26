package gomod

import (
	"reflect"
	"strconv"
)

const (
	EmailTag    string = "email"
	RequiredTag string = "required"
	TypeTag     string = "type"
)

func getValidationRules(m reflect.StructField) (map[string]bool, error) {
	tags := make(map[string]bool)

	if tagType := m.Tag.Get(TypeTag); tagType != "" {
		switch tagType {
		case EmailTag:
			tags[EmailTag] = true
		}
	}

	if tagVal := m.Tag.Get(EmailTag); tagVal != "" {

		boolVal, err := strconv.ParseBool(tagVal)

		if err != nil {
			panic(err)
		}

		if boolVal {
			tags[EmailTag] = boolVal
		}
	}

	if tagVal := m.Tag.Get(RequiredTag); tagVal != "" {

		boolVal, err := strconv.ParseBool(tagVal)

		if err != nil {
			panic(err)
		}

		if boolVal {
			tags[RequiredTag] = boolVal
		}
	}

	return tags, nil
}

func Rules(m reflect.StructField) (map[string]bool, error) {
	return getValidationRules(m)
}
