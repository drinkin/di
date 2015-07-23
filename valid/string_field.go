package valid

import "github.com/drinkin/di/str2"

type StringField struct {
	F     *Field
	Value string
}

func (f *StringField) Enum(vals []string) *StringField {
	if !str2.Slice(vals).Contains(f.Value) {
		f.F.FieldErr("enum")
	}
	return f
}

func (f *StringField) Required() *StringField {
	if f.Value == "" {
		f.F.RequiredErr()
	}
	return f
}
