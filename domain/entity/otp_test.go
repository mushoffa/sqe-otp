package entity

import (
	"strconv"
	"testing"

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