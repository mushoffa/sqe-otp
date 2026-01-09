package domain

import (
	"context"

	"sqe-otp/domain/entity"
)

type OtpRepository interface {
	Insert(context.Context, entity.Otp) error
	FindByCode(context.Context, string) error
	UpdateStatus(context.Context, entity.OtpStatus) error
}
