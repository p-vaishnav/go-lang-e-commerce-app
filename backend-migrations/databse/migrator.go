package database

import (
	"migrations/tables"

	"gorm.io/gorm"
)

type Migrate struct {
	TableName string
	Run       func(*gorm.DB) error
}

func AutoMigrate(db *gorm.DB) []Migrate {
	var user tables.User
	var photo tables.Photo
	var addr tables.Address
	var product tables.Product
	var review tables.Review
	var otpVer tables.OTPVerifications
	var reqResLogs tables.RequestResponseLogs

	userM := Migrate{TableName: "user", Run: func(d *gorm.DB) error { return db.AutoMigrate(&user) }}
	photoM := Migrate{TableName: "photos", Run: func(d *gorm.DB) error { return db.AutoMigrate(&photo) }}
	addrM := Migrate{TableName: "address", Run: func(d *gorm.DB) error { return db.AutoMigrate(&addr) }}
	productM := Migrate{TableName: "products", Run: func(d *gorm.DB) error { return db.AutoMigrate(&product) }}
	reviewM := Migrate{TableName: "reviews", Run: func(d *gorm.DB) error { return db.AutoMigrate(&review) }}
	otpVerM := Migrate{TableName: "otp_verifications", Run: func(d *gorm.DB) error { return db.AutoMigrate(&otpVer) }}
	reqResM := Migrate{TableName: "request_response_logs", Run: func(d *gorm.DB) error { return db.AutoMigrate(&reqResLogs) }}
	// NOTE: table name ain't getting into an effect

	return []Migrate{
		userM,
		photoM,
		addrM,
		productM,
		reviewM,
		otpVerM,
		reqResM,
	}
}
