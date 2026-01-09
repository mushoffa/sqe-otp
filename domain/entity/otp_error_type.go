package entity

var (
	ErrInvalidUserID = OtpError{Message: "invalid_user_id", Description: "User Id Cannot Empty"}
	ErrOtpNotFound   = OtpError{Message: "otp_not_found", Description: "OTP Not Found"}
	ErrOtpInvalid    = OtpError{Message: "otp_invalid", Description: "Invalid OTP"}
	ErrOtpValidated  = OtpError{Message: "otp_validated", Description: "OTP is already validated"}
	ErrOtpExpired    = OtpError{Message: "otp_expired", Description: "OTP is expired"}
)

type OtpError struct {
	Message     string `json:"error"`
	Description string `json:"description"`
}

func (e OtpError) Error() string {
	return e.Message
}
