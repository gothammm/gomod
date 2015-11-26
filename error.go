package gomod

type ModError struct {
	Message   string
	FieldName string
	ModelName string
}

func (e *ModError) String() string {
	return "Model Name: " + e.ModelName + ", Field Name: " + e.FieldName + ", Message: " + e.Message
}
