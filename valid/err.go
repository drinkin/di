package valid

import "fmt"

type FieldErr struct {
	// The name of the struct field with the error.
	Field string `json:"field"`

	// The name of the validation.
	Tag string `json:"tag"`

	// The human readable error.
	Msg string `json:"msg"`
}

func NewFieldError(field, tag string) *FieldErr {
	err := &FieldErr{
		Field: field,
		Tag:   tag,
	}
	err.Msg = err.Error()
	return err
}

// Error implemented to satisfy the error interface
func (err *FieldErr) Error() string {
	return fmt.Sprintf(`Field "%s" failed the "%s" validator.`, err.Field, err.Tag)
}

// StatusCode implemented to satisfy HTTPErr
func (err *FieldErr) StatusCode() int {
	return 400
}
