package entity

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
)

type Otp struct {
	userID string
	code   string
}

func NewOtp(userID string) (Otp, error) {
	if !isUserIdValid(userID) {
		return Otp{}, ErrInvalidUserID
	}
	code, err := generate()
	if err != nil {
		return Otp{}, err
	}
	return Otp{
		userID: userID,
		code:   code,
	}, nil
}

func (e Otp) GetCode() string {
	return e.code
}

func (e Otp) Hash() string {
	payload := fmt.Sprintf("%s%s", e.userID, e.code)
	hasher := sha256.New()
	hasher.Write([]byte(payload))
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
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
	return fmt.Sprintf("%06d",(int(n.Int64()) + min)), nil
}
