package env

import (
	"fmt"
	"os"
)

//MustSet sets the value of the environment variable named by the key, panics on error.
func MustSet(key string, val interface{}) {
	err := os.Setenv(key, fmt.Sprintf("%v", val))
	if err != nil {
		panic(err)
	}
}

//MustGet retrieves the value of the environment variable named by the key. It returns the value, or panics if the variable is not present.
func MustGet(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Errorf("Missing environment variable: %s", key))
	}

	return v
}

//Get retrieves the value of the env var with an optional default.
func Get(key string, d ...interface{}) string {
	v := os.Getenv(key)
	if v == "" && len(d) != 0 {
		return fmt.Sprintf("%v", d[0])
	}
	return v
}
