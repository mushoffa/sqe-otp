package entity

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Otp_Empty_UserID(t *testing.T) {
	// Given
	userID := ""

	// When
	otp, err := NewOtp(userID)

	// Then
	assert.Equal(t, Otp{}, otp)
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidUserID, err)
}

func Test_Otp_NotEmpty_UserID(t *testing.T) {
	// Given
	userID := "john"

	// When
	otp, err := NewOtp(userID)
	code := otp.GetCode()
	_, errCode := strconv.Atoi(code)

	// Then
	assert.NotEqual(t, Otp{}, otp)
	assert.Equal(t, 6, len(code))
	assert.Equal(t, userID, otp.GetUserID())
	assert.Nil(t, err)
	assert.Nil(t, errCode)
}

func Test_Otp_Random_Code_Same_UserID(t *testing.T) {
	// Given
	userID := "john"

	// When
	otp1, err1 := NewOtp(userID)
	otp2, err2 := NewOtp(userID)

	// Then
	assert.NotEqual(t, Otp{}, otp1)
	assert.NotEqual(t, Otp{}, otp2)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.NotEqual(t, otp1.GetCode(), otp2.GetCode())
}

func Test_Otp_Code_Hashed_Not_Empty(t *testing.T) {
	// Given
	userID := "john"

	// When
	otp, err := NewOtp(userID)
	hashed := otp.Hash()

	// Then
	assert.NotEqual(t, Otp{}, otp)
	assert.Nil(t, err)
	assert.NotEqual(t, "", hashed)
}

func Test_Otp_Code_Hashed_Not_Equal(t *testing.T) {
	// Given
	userID := "john"

	// When
	otp, err := NewOtp(userID)
	code := otp.GetCode()
	hashed := otp.Hash()

	// Then
	assert.NotEqual(t, Otp{}, otp)
	assert.Nil(t, err)
	assert.NotEqual(t, code, hashed)
}

func Test_Otp_ValidateStatus_Given_OtpStatusOtpStatusCreated_Then_Return_Nil(t *testing.T) {
	// Given
	otp := Otp{
		Status: OtpStatusCreated,
	}

	// When
	err := otp.ValidateStatus()

	// Then
	assert.Nil(t, err)
}

func Test_Otp_ValidateStatus_Given_OtpStatusValidated_Then_Return_ErrOtpValidated(t *testing.T) {
	// Given
	otp := Otp{
		Status: OtpStatusValidated,
	}

	// When
	err := otp.ValidateStatus()

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, ErrOtpValidated, err)
}

func Test_Otp_ValidateStatus_Given_OtpStatusExpired_Then_Return_ErrOtpExpired(t *testing.T) {
	// Given
	otp := Otp{
		Status: OtpStatusExpired,
	}

	// When
	err := otp.ValidateStatus()

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, ErrOtpExpired, err)
}

func Test_Otp_IsExpired_True(t *testing.T) {
	// Given
	otp := Otp{
		ExpiredAt: time.Now().Add(-time.Duration(3) * time.Minute),
	}

	// When
	isExpired := otp.IsExpired()

	// Then
	assert.True(t, isExpired)
}

func Test_Otp_IsExpired_False(t *testing.T) {
	// Given
	otp := Otp{
		ExpiredAt: time.Now().Add(time.Duration(1) * time.Minute),
	}

	// When
	isExpired := otp.IsExpired()

	// Then
	assert.False(t, isExpired)
}
