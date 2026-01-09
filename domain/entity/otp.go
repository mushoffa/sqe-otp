package entity

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	vo "sqe-otp/domain/valueobject"
)

type Otp struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiredAt time.Time
	UserID    string
	Code      string `gorm:"-"`
	Status    OtpStatus
}

func NewOtp(userID string) (Otp, error) {
	if !isUserIdValid(userID) {
		return Otp{}, ErrInvalidUserID
	}
	code, err := generate()
	if err != nil {
		return Otp{}, err
	}

	now := time.Now()

	return Otp{
		CreatedAt: now,
		UpdatedAt: now,
		ExpiredAt: now.Add(time.Duration(2) * time.Minute),
		UserID:    userID,
		Code:      code,
		Status:    OtpStatusCreated,
	}, nil
}

func (e Otp) GetUserID() string {
	return e.UserID
}

func (e Otp) GetCode() string {
	return e.Code
}

func (e Otp) Hash() string {
	return vo.Hasher(fmt.Sprintf("%s%s", e.UserID, e.Code))
}

func (e Otp) ValidateStatus() error {
	switch e.Status {
	case OtpStatusValidated:
		return ErrOtpValidated
	case OtpStatusExpired:
		return ErrOtpExpired
	default:
		return nil
	}
}

func (e Otp) IsExpired() bool {
	return time.Now().After(e.ExpiredAt)
}

func isUserIdValid(userID string) bool {
	return (userID != "")
}

func generate() (string, error) {
	// Define the minimum value (100000) and the range size (900000).
	// The range will be [0, 899999], so we add the minimum value later.
	min := 100000
	max := 999999
	// The upper bound for rand.Int is exclusive, so we use max - min + 1.
	rangeSize := big.NewInt(int64(max - min + 1))

	// Generate a random number n in the range [0, rangeSize).
	n, err := rand.Int(rand.Reader, rangeSize)
	if err != nil {
		return "", err
	}

	// Add the minimum value to the generated number to get a number in the desired range [100000, 999999].
	return fmt.Sprintf("%06d", (int(n.Int64()) + min)), nil
}
