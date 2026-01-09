package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type PostgresClientMock struct {
	mock.Mock
}

func (m *PostgresClientMock) Insert(ctx context.Context, table any) error {
	args := m.Called(ctx, table)
	return args.Error(0)
}

func (m *PostgresClientMock) QueryByCondition(ctx context.Context, conditions map[string]any, table any) error {
	args := m.Called(ctx, conditions, table)
	return args.Error(0)
}

func (m *PostgresClientMock) UpdateByCondition(ctx context.Context, conditions map[string]any, table any) error {
	args := m.Called(ctx, conditions, table)
	return args.Error(0)
}

func (m *PostgresClientMock) Shutdown() {

}
