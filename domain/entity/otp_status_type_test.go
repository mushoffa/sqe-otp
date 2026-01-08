package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OtpStatusCreated(t *testing.T) {
	// Given
	otp := OtpStatusCreated

	// When
	status := otp.String()

	// Then
	assert.Equal(t, "created", status)
}

func Test_OtpStatusValidated(t *testing.T) {
	// Given
	otp := OtpStatusValidated

	// When
	status := otp.String()

	// Then
	assert.Equal(t, "validated", status)
}

func Test_OtpStatusExpired(t *testing.T) {
	// Given
	otp := OtpStatusExpired

	// When
	status := otp.String()

	// Then
	assert.Equal(t, "expired", status)
}
