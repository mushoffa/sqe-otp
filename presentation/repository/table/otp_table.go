package table

import (
	"sqe-otp/domain/entity"

	"gorm.io/gorm"
)

type OtpTable struct{}

func (db OtpTable) TableName() string {
	return "otps"
}

type InsertOtp struct {
	OtpTable
	Otp  entity.Otp `gorm:"embedded"`
	Code string
}

func (db *InsertOtp) BeforeCreate(tx *gorm.DB) (err error) {
	db.Code = db.Otp.Hash()
	return
}

type QueryOtp struct {
	OtpTable
	Otp  entity.Otp `gorm:"embedded"`
	Code string
}

type UpdateOtpStatus struct {
	OtpTable
	Status entity.OtpStatus
}
