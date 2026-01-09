package usecase

import (
	"context"
	"fmt"

	"sqe-otp/domain/entity"
	"sqe-otp/domain/repository"
	vo "sqe-otp/domain/valueobject"
)

type OtpUsecase interface {
	GenerateOtp(context.Context, string) (string, error)
	ValidateOtp(context.Context, string, string) error
}

type otp struct {
	r domain.OtpRepository
}

func NewOtpUsecase(r domain.OtpRepository) OtpUsecase {
	return &otp{
		r: r,
	}
}

func (u *otp) GenerateOtp(ctx context.Context, userID string) (string, error) {
	otp, err := entity.NewOtp(userID)
	if err != nil {
		return "", err
	}

	if err := u.r.Insert(ctx, otp); err != nil {
		return "", err
	}

	return otp.GetCode(), nil
}

func (u *otp) ValidateOtp(ctx context.Context, userID, code string) error {
	payload := fmt.Sprintf("%s%s", userID, code)
	hashedOtp := vo.Hasher(payload)
	otp, err := u.r.FindByCode(ctx, hashedOtp)
	if err != nil {
		return err
	}

	if err := otp.ValidateStatus(); err != nil {
		return err
	}

	if otp.IsExpired() {
		u.r.UpdateStatus(ctx, hashedOtp, entity.OtpStatusExpired)
		return entity.ErrOtpExpired
	}

	u.r.UpdateStatus(ctx, hashedOtp, entity.OtpStatusValidated)
	return nil
}
