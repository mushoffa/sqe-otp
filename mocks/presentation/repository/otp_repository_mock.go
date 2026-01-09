package mocks

import (
	"context"

	"sqe-otp/domain/entity"

	"github.com/stretchr/testify/mock"
)

type OtpRepositoryMock struct {
	mock.Mock
}

func (m *OtpRepositoryMock) Insert(ctx context.Context, otp entity.Otp) error {
	args := m.Called(ctx, otp)
	return args.Error(0)
}

func (m *OtpRepositoryMock) FindByCode(ctx context.Context, code string) (entity.Otp, error) {
	args := m.Called(ctx, code)
	return args.Get(0).(entity.Otp), args.Error(1)
}

func (m *OtpRepositoryMock) UpdateStatus(ctx context.Context, code string, status entity.OtpStatus) error {
	args := m.Called(ctx, code, status)
	return args.Error(0)
}
