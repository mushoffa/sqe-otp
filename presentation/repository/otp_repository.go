package repository

import (
	"context"
	"errors"

	"sqe-otp/domain/entity"
	"sqe-otp/domain/repository"
	"sqe-otp/infrastructure/postgres"
	"sqe-otp/presentation/repository/table"

	"gorm.io/gorm"
)

type otp struct {
	db postgres.DatabaseService
}

func NewOtpRepository(db postgres.DatabaseService) domain.OtpRepository {
	return &otp{
		db: db,
	}
}

func (r *otp) Insert(ctx context.Context, otp entity.Otp) error {
	return r.db.Insert(ctx, &table.InsertOtp{
		Otp: otp,
	})
}

func (r *otp) FindByCode(ctx context.Context, code string) (entity.Otp, error) {
	query := table.QueryOtp{}

	if err := r.db.QueryByCondition(ctx, map[string]any{
		"code": code,
	}, &query); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Otp{}, entity.ErrOtpNotFound
		}
		return entity.Otp{}, entity.OtpError{Message: "internal_server_error", Description: err.Error()}
	}

	otp := query.Otp
	otp.Code = query.Code
	return otp, nil
}

func (r *otp) UpdateStatus(ctx context.Context, code string, status entity.OtpStatus) error {
	return r.db.UpdateByCondition(ctx, map[string]any{
		"code": code,
	}, &table.UpdateOtpStatus{Status: status})
}
