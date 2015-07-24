package gvalid

import (
	"fmt"
	"reflect"

	"github.com/drinkin/di/valid"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

func MatchEnumErr(field string) types.GomegaMatcher {
	return MatchFieldErr(field, "enum")
}

func MatchRequiredErr(field string) types.GomegaMatcher {
	return MatchFieldErr(field, "required")
}

func MatchFieldErr(field, tag string) types.GomegaMatcher {
	return &MatchFieldErrMatcher{
		Field:    field,
		Tag:      tag,
		Expected: valid.NewFieldError(field, tag),
	}
}

type MatchFieldErrMatcher struct {
	Field    string
	Tag      string
	Expected *valid.FieldErr
}

func (m *MatchFieldErrMatcher) Match(a interface{}) (success bool, err error) {
	if a == nil {
		return false, fmt.Errorf("Expected an error, got nil")
	}

	actualFieldErr, ok := a.(*valid.FieldErr)
	if !ok {
		return false, fmt.Errorf("Expected a FieldError.  Got:\n%s", format.Object(a, 1))
	}

	return reflect.DeepEqual(actualFieldErr, m.Expected), nil
}

func (matcher *MatchFieldErrMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to match FieldError", matcher.Expected)
}

func (matcher *MatchFieldErrMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to match FieldError", matcher.Expected)
}
