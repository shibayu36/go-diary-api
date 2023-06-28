package model

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateApiKey(t *testing.T) {
	apiKey, err := GenerateApiKey()
	assert.Nil(t, err)
	assert.Regexp(t, regexp.MustCompile("^[0-9a-f]{64}$"), apiKey)
}
