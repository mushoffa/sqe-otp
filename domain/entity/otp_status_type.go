package entity

import (
	"database/sql/driver"
	"fmt"
)

type OtpStatus int

const (
	OtpStatusCreated OtpStatus = iota
	OtpStatusValidated
	OtpStatusExpired
)

func (e OtpStatus) String() string {
	return [...]string{"created", "validated", "expired"}[e]
}

func (e *OtpStatus) Scan(value any) error {
	status, ok := value.(string)
	if !ok {
		return fmt.Errorf("Cannot Scan Type %T OTP Status", value)
	}

	switch status {
	case "created":
		*e = OtpStatusCreated
	case "validated":
		*e = OtpStatusValidated
	case "expired":
		*e = OtpStatusExpired
	default:
		return fmt.Errorf("Unknown OTP Status Type %T", value)
	}
	return nil
}

func (e OtpStatus) Value() (driver.Value, error) {
	return e.String(), nil
}
