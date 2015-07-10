package random_test

import (
	"testing"

	"github.com/drinkin/di/random"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert := assert.New(t)
	assert.NotEqual(random.String(10), random.String(10))
}

func TestBase64(t *testing.T) {
	assert := assert.New(t)
	assert.NotEqual(random.Base64(32), random.Base64(32))
}
