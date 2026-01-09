package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type OtpUsecaseMock struct {
	mock.Mock
}

func (m *OtpUsecaseMock) GenerateOtp(ctx context.Context, userID string) (string, error) {
	args := m.Called(ctx, userID)
	return args.String(0), args.Error(1)
}

func (m *OtpUsecaseMock) ValidateOtp(ctx context.Context, userID, code string) error {
	args := m.Called(ctx, userID, code)
	return args.Error(0)
}
