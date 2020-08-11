package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	encoder := NewPasswordSerice()
	encodedString := encoder.Encode("secret")
	t.Log(encodedString)
	assert.NotEmpty(t, encodedString)

}

func TestMatches(t *testing.T) {
	encoder := NewPasswordSerice()
	hash := encoder.Encode("secret")
	assert.True(t, encoder.Matches("secret", hash))
}
