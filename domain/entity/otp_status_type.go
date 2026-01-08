package entity

type OtpStatus int

const (
	OtpStatusCreated OtpStatus = iota
	OtpStatusValidated
	OtpStatusExpired
)

func (e OtpStatus) String() string {
	return [...]string{"created", "validated", "expired"}[e]
}
