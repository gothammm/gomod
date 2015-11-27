package gomod

import (
	"testing"
)

type TestModel struct {
	Email string `json:"name" email:"true"`
	Age   int    `json:"age" max:"10"`
	Phone int64  `json:"phone" required:"true"`
}

func TestIsStruct(t *testing.T) {
	t.Parallel()

	if IsStruct(1) {
		t.Error("Value should have a struct value")
	}

	if !IsStruct(&TestModel{Email: "test@atest.com", Age: 20}) {
		t.Error("IsStruct must return true")
	}

	if !IsStruct(TestModel{}) {
		t.Error("IsStruct must return true")
	}
}

func TestValidate(t *testing.T) {
	t.Parallel()

	modOne := &TestModel{Email: "test@test.com", Age: 10, Phone: 9885239317}
	var errorOne Errors
	errorOne, err := Validate(modOne)

	if err != nil {
		t.Error(err)
	}

	if errorOne != nil {
		t.Error("Expected no model errors but got", len(errorOne), "model error(s) instead")
		for _, j := range errorOne.Json() {
			t.Error(j.String())
		}
	}

	t.Log("-------------------------------------")
	modTwo := &TestModel{Email: "testsg.com", Age: 10}

	var errorTwo Errors

	errorTwo, errt := Validate(modTwo)

	if errt != nil {
		t.Error(errt)
	}
	for _, j := range errorTwo.Json() {
		t.Log(j.String())
	}
	if errorTwo == nil {
		t.Error("Expected model errors, but got 0 model errors instead")
	}
}
