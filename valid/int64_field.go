package valid

type Int64Field struct {
	F     *Field
	Value int64
}

func (f *Int64Field) Equal(v int64) *Int64Field {
	if f.Value != v {
		f.F.FieldErr("equal")
	}
	return f
}

func (f *Int64Field) Required() *Int64Field {
	if f.Value == 0 {
		f.F.RequiredErr()
	}

	return f
}
