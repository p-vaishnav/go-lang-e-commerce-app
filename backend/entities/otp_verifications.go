package entities

import "time"

type OTPVerifications struct {
	ID          int    `gorm:"column:otp_verifications_id;primaryKey;autoIncrement"`
	PID         string `gorm:"column:otp_verifications_pid;unique;not null;type:varchar(40)"`
	Purpose     string `gorm:"column:otp_verifications_purpose;type:varchar(25)"`
	Type        string `gorm:"column:type;type:varchar(25)"`
	Medium      string `gorm:"column:medium;type:varchar(25)"`
	OTP         string `gorm:"column:otp;type:varchar(8)"`
	Provider    string `gorm:"column:provider;type:varchar(25)"`
	Status      string `gorm:"column:status;type:varchar(20)"`
	Mobile      string `gorm:"column:mobile;type:varchar(20)"`
	Email       string `gorm:"column:email;type:varchar(80)"`
	CountryCode string `gorm:"column:country_code;type:varchar(10)"`
	OTPCount    int    `gorm:"column:otp_count"`
	RetryCount  int    `gorm:"column:retry_count"`
	IsSandbox   bool   `gorm:"column:is_sandbox;default:false"`
	IsDeleted   bool   `gorm:"column:is_deleted;default:false"`
	IsActive    bool   `gorm:"column:is_active;default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
