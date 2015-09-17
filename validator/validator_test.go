package validator

import (
	"testing"

	"github.com/samora/ussd-go/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestLength(t *testing.T) {
	assert := assert.New(t)
	f := Map["length"]

	err := f("Full Name", "Samora", "10")
	assert.Error(err)

	err = f("Full Name", "Samora Dake", "10")
	assert.Nil(err)

	err = f("PIN", "1234", "4", "4")
	assert.Nil(err)

	err = f("PIN", "12345", "4", "4")
	assert.Error(err)
}

func TestRange(t *testing.T) {
	assert := assert.New(t)
	f := Map["range"]

	err := f("Age", "15", "18", "35")
	assert.Error(err)

	err = f("Age", "18", "18", "35")
	assert.Nil(err)
}
