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

	userM := Migrate{TableName: "user", Run: func(d *gorm.DB) error { return db.AutoMigrate(&user) }}
	photoM := Migrate{TableName: "photos", Run: func(d *gorm.DB) error { return db.AutoMigrate(&photo) }}
	addrM := Migrate{TableName: "address", Run: func(d *gorm.DB) error { return db.AutoMigrate(&addr) }}
	productM := Migrate{TableName: "products", Run: func(d *gorm.DB) error { return db.AutoMigrate(&product) }}
	reviewM := Migrate{TableName: "reviews", Run: func(d *gorm.DB) error { return db.AutoMigrate(&review) }}
	otpVerM := Migrate{TableName: "otp_verifications", Run: func(d *gorm.DB) error { return db.AutoMigrate(&otpVer) }}

	return []Migrate{
		userM,
		photoM,
		addrM,
		productM,
		reviewM,
		otpVerM,
	}
}
