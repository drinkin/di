package valid

type Field struct {
	Name string

	errs []*FieldErr
}

func (f *Field) String(v string) *StringField {
	return &StringField{F: f, Value: v}
}

func (f *Field) Int64(v int64) *Int64Field {
	return &Int64Field{F: f, Value: v}
}

func (f *Field) FieldErr(errName string) {
	f.errs = append(f.errs, NewFieldError(f.Name, errName))
}

func (f *Field) RequiredErr() {
	f.FieldErr("required")
}

func (f *Field) Check() error {
	for _, err := range f.errs {
		if err != nil {
			return err
		}
	}
	return nil
}
