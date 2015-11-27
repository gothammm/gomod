package gomod

import (
	"reflect"
	"strconv"
)

const (
	EmailTag      string = "email"
	RequiredTag   string = "required"
	TypeTag       string = "type"
	PhoneIndiaTag string = "phone_IN"
	NumericTag    string = "numeric"
	MaxLenTag     string = "max"
	MinLenTag     string = "min"
)

func getValidationRules(m reflect.StructField) (map[string]interface{}, error) {
	tags := make(map[string]interface{})

	if tagType := m.Tag.Get(TypeTag); tagType != "" {
		switch tagType {
		case EmailTag:
			tags[EmailTag] = true
		case PhoneIndiaTag:
			tags[PhoneIndiaTag] = true
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

	if tagVal := m.Tag.Get(PhoneIndiaTag); tagVal != "" {
		boolVal, err := strconv.ParseBool(tagVal)
		if err != nil {
			panic(err)
		}
		if boolVal {
			tags[PhoneIndiaTag] = boolVal
		}
	}

	if tagVal := m.Tag.Get(MaxLenTag); tagVal != "" {
		maxLen, err := strconv.Atoi(tagVal)

		if err == nil {
			tags[MaxLenTag] = maxLen
		}
	}

	if tagVal := m.Tag.Get(MinLenTag); tagVal != "" {
		minLen, err := strconv.Atoi(tagVal)

		if err == nil {
			tags[MinLenTag] = minLen
		}
	}

	return tags, nil
}

func Rules(m reflect.StructField) (map[string]interface{}, error) {
	return getValidationRules(m)
}
