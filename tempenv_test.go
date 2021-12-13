package tempenv_test

import (
	"github.com/jdockerty/tempenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	assert := assert.New(t)
	expected := "Hello, World!"
	assert.Equal(expected, tempenv.Hello())
}
