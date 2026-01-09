package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OtpError_Error(t *testing.T) {
	// Given
	message := "internal_server_error"
	description := "description"

	// When
	err := OtpError{
		Message:     message,
		Description: description,
	}

	// Then
	assert.Equal(t, message, err.Error())
}
