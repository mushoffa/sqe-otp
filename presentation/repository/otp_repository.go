package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"sqe-otp/domain/entity"
	"sqe-otp/domain/repository"
	"sqe-otp/infrastructure/postgres"
	"sqe-otp/infrastructure/redis"
	"sqe-otp/presentation/repository/table"

	"gorm.io/gorm"
)

var (
	redisOtpKey = "redis:otp:%s"
)

type otp struct {
	db    postgres.DatabaseService
	redis redis.RedisService
}

func NewOtpRepository(db postgres.DatabaseService, redis redis.RedisService) domain.OtpRepository {
	return &otp{
		db:    db,
		redis: redis,
	}
}

func (r *otp) Insert(ctx context.Context, otp entity.Otp) error {
	redisKey := getRedisKey(otp.UserID)
	isExist, err := r.redis.Exists(ctx, redisKey)
	if err != nil {
		return err
	}

	if isExist {
		return entity.ErrOtpRequested
	}

	if err := r.db.Insert(ctx, &table.InsertOtp{
		Otp: otp,
	}); err != nil {
		return err
	}

	// Not good, just a temporary solution
	r.StoreSession(ctx, otp.UserID)
	return nil
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

func (r *otp) StoreSession(ctx context.Context, userID string) {
	r.redis.Set(ctx, getRedisKey(userID), true, time.Duration(2)*time.Minute)
}

func (r *otp) ClearSession(ctx context.Context, userID string) {
	r.redis.Del(ctx, getRedisKey(userID))
}

func getRedisKey(userID string) string {
	return fmt.Sprintf("%s%s", redisOtpKey, userID)
}
