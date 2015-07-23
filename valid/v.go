package valid

import "github.com/drinkin/di/str2"

type V struct {
	fields []*Field
}

func New() *V {
	return &V{}
}

func (v *V) F(name string) *Field {
	f := &Field{Name: name}
	v.fields = append(v.fields, f)
	return f
}

// Check returns an error for the first validation failed.
// if `names` are provided, only fields with the given name are checked.
func (v *V) Check(names ...string) error {
	n := str2.Slice(names)

	for _, f := range v.fields {

		if len(names) == 0 || n.Contains(f.Name) {
			if err := f.Check(); err != nil {
				return err
			}
		}

	}
	return nil
}
