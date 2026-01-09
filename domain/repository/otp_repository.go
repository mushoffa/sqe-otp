package domain

import (
	"context"

	"sqe-otp/domain/entity"
)

type OtpRepository interface {
	Insert(context.Context, entity.Otp) error
	FindByCode(context.Context, string) (entity.Otp, error)
	UpdateStatus(context.Context, string, entity.OtpStatus) error
	StoreSession(context.Context, string)
	ClearSession(context.Context, string)
}
