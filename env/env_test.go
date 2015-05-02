package env_test

import (
	"testing"

	"github.com/drinkin/di/env"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	key := "TEST_ENV_VAR"
	value := "testrandom"
	defVal := "default"
	env.MustSet(key, value)

	assert.Equal(t, value, env.Get(key))
	assert.Equal(t, value, env.Get(key, defVal))
	assert.Equal(t, "", env.Get("NON_EXISTANT_ENV_VAR"))
	assert.Equal(t, defVal, env.Get("NON_EXISTANT_ENV_VAR", defVal))
}

func TestMustGet(t *testing.T) {
	key := "TEST_ENV_MUST_GET"
	value := "testrandom"

	env.MustSet(key, value)
	assert.Equal(t, value, env.MustGet(key))

	assert.Panics(t, func() {
		env.MustGet("NON_EXISTANT_ENV_VAR")
	}, "Should panic for non-existant env var")
}
