package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Hasher(t *testing.T) {
	// Given
	payload := ""

	// When
	hashed := Hasher(payload)

	// Then
	assert.NotEmpty(t, hashed)
}
