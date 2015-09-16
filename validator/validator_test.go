package validator

import (
	"testing"

	"github.com/samora/ussd-go/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestLength(t *testing.T) {
	assert := assert.New(t)

	err := Length("Full Name", "Samora", "10")
	assert.Error(err)

	err = Length("Full Name", "Samora Dake", "10")
	assert.Nil(err)

	err = Length("PIN", "1234", "4", "4")
	assert.Nil(err)

	err = Length("PIN", "12345", "4", "4")
	assert.Error(err)
}
