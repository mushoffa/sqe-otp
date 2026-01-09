package repository

import (
	"context"
	"testing"

	"sqe-otp/domain/entity"
	mocks "sqe-otp/mocks/infrastructure/postgres"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func Test_Insert_Given_Return_Success(t *testing.T) {
	mockPostgres := new(mocks.PostgresClientMock)
	r := NewOtpRepository(mockPostgres)

	// Given
	otp := entity.Otp{}
	mockPostgres.On("Insert", mock.Anything, mock.Anything).Return(nil)

	// When
	err := r.Insert(context.Background(), otp)

	// Then
	assert.Nil(t, err)
}

func Test_Insert_Given_Return_Error(t *testing.T) {
	mockPostgres := new(mocks.PostgresClientMock)
	r := NewOtpRepository(mockPostgres)

	// Given
	otp := entity.Otp{}
	mockPostgres.On("Insert", mock.Anything, mock.Anything).Return(gorm.ErrNotImplemented)

	// When
	err := r.Insert(context.Background(), otp)

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrNotImplemented, err)
}

func Test_FindByCode_Given_OtpQuery_Then_Return_Success(t *testing.T) {
	mockPostgres := new(mocks.PostgresClientMock)
	r := NewOtpRepository(mockPostgres)

	// Given
	code := "123456"
	mockPostgres.On("QueryByCondition", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// When
	otp, err := r.FindByCode(context.Background(), code)

	// Then
	assert.Equal(t, entity.Otp{}, otp)
	assert.Nil(t, err)
}

func Test_FindByCode_Given_NotFound_Then_Return_ErrOtpNotFound(t *testing.T) {
	mockPostgres := new(mocks.PostgresClientMock)
	r := NewOtpRepository(mockPostgres)

	// Given
	code := "123456"
	mockPostgres.On("QueryByCondition", mock.Anything, mock.Anything, mock.Anything).Return(gorm.ErrRecordNotFound)

	// When
	otp, err := r.FindByCode(context.Background(), code)

	// Then
	assert.Equal(t, entity.Otp{}, otp)
	assert.NotNil(t, err)
	assert.Equal(t, entity.ErrOtpNotFound, err)
}

func Test_FindByCode_Given_ErrInvalidValue_Then_Return_Error(t *testing.T) {
	mockPostgres := new(mocks.PostgresClientMock)
	r := NewOtpRepository(mockPostgres)

	// Given
	code := "123456"
	mockPostgres.On("QueryByCondition", mock.Anything, mock.Anything, mock.Anything).Return(gorm.ErrInvalidValue)

	// When
	otp, err := r.FindByCode(context.Background(), code)

	// Then
	assert.Equal(t, entity.Otp{}, otp)
	assert.NotNil(t, err)
}

func Test_UpdateByCondition_Given_Success_Update_Then_Return_Nil(t *testing.T) {
	mockPostgres := new(mocks.PostgresClientMock)
	r := NewOtpRepository(mockPostgres)

	// Given
	code := "123456"
	mockPostgres.On("UpdateByCondition", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// When
	err := r.UpdateStatus(context.Background(), code, entity.OtpStatusValidated)

	// Then
	assert.Nil(t, err)
}
