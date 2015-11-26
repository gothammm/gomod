package gomod

import (
	"testing"
)

type TestModel struct {
	Email string `json:"name" email:"true" required:"true"`
	Age   int    `json:"age" required:"true"`
}

func TestIsStruct(t *testing.T) {
	t.Parallel()

	if IsStruct(1) {
		t.Error("Value should have a struct value")
	}

	if !IsStruct(&TestModel{Email: "test@test.com", Age: 10}) {
		t.Error("IsStruct must return true")
	}

	if !IsStruct(TestModel{}) {
		t.Error("IsStruct must return true")
	}
}

func TestValidate(t *testing.T) {
	t.Parallel()

	modOne := &TestModel{Email: "test@test.com", Age: 10}

	errorOne, err := Validate(modOne)

	if err != nil {
		t.Error(err)
	}

	if errorOne != nil {
		t.Error("Expected no model errors but got", len(errorOne), "model error(s) instead")
	}

	modTwo := TestModel{Email: "testsg.com", Age: 10}

	errTwo, err := Validate(modTwo)

	if err != nil {
		t.Error(err)
	}

	if errTwo == nil {
		t.Error("Expected model errors, but got 0 model errors instead")
	}
}
