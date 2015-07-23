package valid

type Int64Field struct {
	F     *Field
	Value int64
}

func (f *Int64Field) Required() *Int64Field {
	if f.Value == 0 {
		f.F.RequiredErr()
	}

	return f
}
