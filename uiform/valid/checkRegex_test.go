package valid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckRegex(t *testing.T) {
	m, err := CheckRegex("^[a-z]+$", "abc")
	assert.Nil(t, err)
	assert.Equal(t, true, m)
}
