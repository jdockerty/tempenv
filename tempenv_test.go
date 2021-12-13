package tempenv_test

import (
	"github.com/jdockerty/tempenv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	testTempEnv = "TEMPENV_TEST"
	during      = "duringAction"
	after       = "afterAction"
)

// TestSetEnvValidateFunc asserts whether our test function, which returns the current environment variable, has
// the appropriate variable set during its execution time.
func TestSetEnvValidateFunc(t *testing.T) {
	assert := assert.New(t)

	testerFunc := func(envVar string) string {
		return os.Getenv(envVar)
	}

	funcArgument := testTempEnv
	result, err := tempenv.Set(testerFunc, testTempEnv, during, after, funcArgument)
	assert.Nil(err)

	actual := (*result)[0].String() // We know that our function returns a string
	assert.Equalf(during, actual, "Expected %s string, got %s\n", during, actual)
	assert.Equal(after, os.Getenv(testTempEnv), "Expected %s, got %s\n", after, actual)
}
