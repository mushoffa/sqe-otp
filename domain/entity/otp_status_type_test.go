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

func Test_OtpStatus_Scan(t *testing.T) {
	// Given
	value := 123
	status := new(OtpStatus)

	// When
	err := status.Scan(value)

	// Then
	assert.NotNil(t, err)
}

func Test_OtpStatus_Scan_Given_created_Then_Return_OtpStatusCreated(t *testing.T) {
	// Given
	value := "created"
	status := new(OtpStatus)

	// When
	err := status.Scan(value)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, OtpStatusCreated, *status)
}

func Test_OtpStatus_Scan_Given_validated_Then_Return_OtpStatusValidated(t *testing.T) {
	// Given
	value := "validated"
	status := new(OtpStatus)

	// When
	err := status.Scan(value)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, OtpStatusValidated, *status)
}

func Test_OtpStatus_Scan_Given_expired_Then_Return_OtpStatusExpired(t *testing.T) {
	// Given
	value := "expired"
	status := new(OtpStatus)

	// When
	err := status.Scan(value)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, OtpStatusExpired, *status)
}

func Test_OtpStatus_Scan_Given_unknown_status_Then_Return_Error(t *testing.T) {
	// Given
	value := "deleted"
	status := new(OtpStatus)

	// When
	err := status.Scan(value)

	// Then
	assert.NotNil(t, err)
	assert.NotEqual(t, OtpStatusCreated, err)
	assert.NotEqual(t, OtpStatusValidated, err)
	assert.NotEqual(t, OtpStatusExpired, err)
}

func Test_OtpStatus_Value_OtpStatusCreated(t *testing.T) {
	// Given
	status := OtpStatusCreated

	// When
	value, err := status.Value()

	// Then
	assert.Equal(t, "created", value)
	assert.Nil(t, err)
}

func Test_OtpStatus_Value_OtpStatusValidated(t *testing.T) {
	// Given
	status := OtpStatusValidated

	// When
	value, err := status.Value()

	// Then
	assert.Equal(t, "validated", value)
	assert.Nil(t, err)
}

func Test_OtpStatus_Value_OtpStatusExpired(t *testing.T) {
	// Given
	status := OtpStatusExpired

	// When
	value, err := status.Value()

	// Then
	assert.Equal(t, "expired", value)
	assert.Nil(t, err)
}
