package usecase

import (
	"context"
	"testing"
	"time"

	"sqe-otp/domain/entity"

	mocks "sqe-otp/mocks/presentation/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GenerateOtp_Given_Success_Insert(t *testing.T) {
	mockRepository := new(mocks.OtpRepositoryMock)
	u := NewOtpUsecase(mockRepository)

	// Given
	userID := "John"
	mockRepository.On("Insert", mock.Anything, mock.Anything).Return(nil)

	// When
	otp, err := u.GenerateOtp(context.Background(), userID)

	// Then
	assert.Equal(t, 6, len(otp))
	assert.Nil(t, err)
}

func Test_GenerateOtp_Given_Empty_User_Then_Return_ErrInvalidUserID(t *testing.T) {
	mockRepository := new(mocks.OtpRepositoryMock)
	u := NewOtpUsecase(mockRepository)

	// Given
	userID := ""

	// When
	otp, err := u.GenerateOtp(context.Background(), userID)

	// Then
	assert.Empty(t, otp)
	assert.NotNil(t, err)
	assert.Equal(t, entity.ErrInvalidUserID, err)
}

func Test_ValidateOtp_Given_Otp_NotFound_Then_Return_ErrOtpNotFound(t *testing.T) {
	mockRepository := new(mocks.OtpRepositoryMock)
	u := NewOtpUsecase(mockRepository)

	// Given
	userID := "John"
	code := "1234"
	mockRepository.On("FindByCode", mock.Anything, mock.Anything).Return(entity.Otp{}, entity.ErrOtpNotFound)

	// When
	err := u.ValidateOtp(context.Background(), userID, code)

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, entity.ErrOtpNotFound, err)
}

func Test_ValidateOtp_Given_Otp_OtpStatusValidated_Then_Return_ErrOtpValidated(t *testing.T) {
	mockRepository := new(mocks.OtpRepositoryMock)
	u := NewOtpUsecase(mockRepository)

	// Given
	userID := "John"
	code := "1234"
	mockOtp := entity.Otp{
		Status: entity.OtpStatusValidated,
	}
	mockRepository.On("FindByCode", mock.Anything, mock.Anything).Return(mockOtp, nil)

	// When
	err := u.ValidateOtp(context.Background(), userID, code)

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, entity.ErrOtpValidated, err)
}

func Test_ValidateOtp_Given_Otp_OtpStatusExpired_Then_Return_ErrOtpExpired(t *testing.T) {
	mockRepository := new(mocks.OtpRepositoryMock)
	u := NewOtpUsecase(mockRepository)

	// Given
	userID := "John"
	code := "1234"
	mockOtp := entity.Otp{
		Status: entity.OtpStatusExpired,
	}
	mockRepository.On("FindByCode", mock.Anything, mock.Anything).Return(mockOtp, nil)

	// When
	err := u.ValidateOtp(context.Background(), userID, code)

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, entity.ErrOtpExpired, err)
}

func Test_ValidateOtp_Given_Valid_Otp_Then_Return_Nil(t *testing.T) {
	mockRepository := new(mocks.OtpRepositoryMock)
	u := NewOtpUsecase(mockRepository)

	// Given
	userID := "John"
	code := "1234"
	mockOtp := entity.Otp{
		ExpiredAt: time.Now().Add(time.Duration(2) * time.Minute),
	}
	mockRepository.On("FindByCode", mock.Anything, mock.Anything).Return(mockOtp, nil)
	mockRepository.On("UpdateStatus", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// When
	err := u.ValidateOtp(context.Background(), userID, code)

	// Then
	assert.Nil(t, err)
}

func Test_ValidateOtp_Given_Expired_Otp_Then_Return_ErrOtpExpired(t *testing.T) {
	mockRepository := new(mocks.OtpRepositoryMock)
	u := NewOtpUsecase(mockRepository)

	// Given
	userID := "John"
	code := "1234"
	mockOtp := entity.Otp{
		ExpiredAt: time.Now().Add(-time.Duration(3) * time.Minute),
	}
	mockRepository.On("FindByCode", mock.Anything, mock.Anything).Return(mockOtp, nil)
	mockRepository.On("UpdateStatus", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// When
	err := u.ValidateOtp(context.Background(), userID, code)

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, entity.ErrOtpExpired, err)
}
