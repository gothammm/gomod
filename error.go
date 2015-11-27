package gomod

type ModError struct {
	Message string
	Field   *Field
}

type JsonError struct {
	Message string
	Field string
}

type Errors []*ModError

func (e *ModError) String() string {
	return "Field: " + e.Field.Name + " Message: " + e.Message
}

func (e Errors) Json() []JsonError {
	var jsonErrors []JsonError
	for _, j := range e {
		name := j.Field.JsonName
		if name == "" {
			name = j.Field.Name
		}
		jsonErrors = append(jsonErrors, JsonError{ Message: j.Message, Field: name })
	}
	return jsonErrors
}

func (j JsonError) String() string {
	return "Field: " + j.Field + " Message: " + j.Message
}
