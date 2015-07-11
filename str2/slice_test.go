package str2_test

import (
	"testing"

	"github.com/drinkin/di/str2"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	assert := assert.New(t)

	s := str2.Slice{"a", "b", "c"}

	for _, v := range s {
		assert.True(s.Contains(v))
	}

	assert.False(s.Contains("z"))

}
