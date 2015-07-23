package valid_test

import (
	"testing"

	"github.com/drinkin/di/valid"
	"github.com/stretchr/testify/assert"
)

func TestInt64Required(t *testing.T) {
	assert := assert.New(t)
	v := valid.New()

	v.F("id").Int64(1).Required()
	assert.NoError(v.Check())

	v.F("id").Int64(0).Required()
	err := v.Check()
	assert.Error(err)
	ferr := err.(*valid.FieldErr)
	assert.Equal(ferr.Field, "id")
	assert.Equal(ferr.Tag, "required")
	assert.Equal(ferr.StatusCode(), 400)
}

func TestStringRequired(t *testing.T) {
	assert := assert.New(t)
	v := valid.New()

	v.F("id").String("hi").Required()
	assert.NoError(v.Check())

	v.F("id").String("").Required()
	err := v.Check()
	assert.Error(err)
	ferr := err.(*valid.FieldErr)
	assert.Equal(ferr.Field, "id")
	assert.Equal(ferr.Tag, "required")
	assert.Equal(ferr.StatusCode(), 400)
}

func TestStringEnum(t *testing.T) {
	assert := assert.New(t)
	values := []string{"a", "b"}

	for _, s := range values {
		v := valid.New()
		v.F("valid_id").String(s).Enum(values)
		assert.NoError(v.Check())
	}

	v := valid.New()
	v.F("valid_id").String("c").Enum(values)
	err := v.Check()
	assert.Error(err)
	ferr := err.(*valid.FieldErr)
	assert.Equal(ferr.Field, "valid_id")
	assert.Equal(ferr.Tag, "enum")
	assert.Equal(ferr.StatusCode(), 400)
}
