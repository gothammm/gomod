package gomod

type ModError struct {
	Message string
	Field   *Field
}

func (e *ModError) String() string {
	return "Field: " + e.Field.Name + " Message: " + e.Message
}
