package command

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestConfig(t *testing.T) {
	action := "Building"

	args := "build"

	config := NewConfig("Building", "build")

	assert.Equal(t, action, config.Action)
	assert.Equal(t, args, config.Args)
}
