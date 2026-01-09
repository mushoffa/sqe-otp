package controller

import (
	"testing"

	mocks "sqe-otp/mocks/usecase"

	"github.com/gofiber/fiber/v2"
)

func Test_CreateOtp_(t *testing.T) {
	// Given
	mockUsecase := new(mocks.OtpUsecaseMock)
	controller := NewController(mockUsecase)

	// When

	// Then
}
