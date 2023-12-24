package tables

import "time"

type Product struct {
	ID            int     `gorm:"column:product_id;primaryKey;autoIncrement;unique;not null;"`
	ProductPID    string  `gorm:"column:product_pid;type: varchar(50);not null;unique;"`
	UserPID       string  `gorm:"column:user_pid;type: varchar(50);not null;unique;"`
	Name          string  `gorm:"column:product_name;type: varchar(50);"`
	Price         float64 `gorm:"column:product_price;type: float64;"` // ??
	Category      string  `gorm:"column:product_category;type: varchar(50);"`
	Stock         int     `gorm:"column:product_stock;"`
	Rating        float64 `gorm:"column:product_rating;"` //
	NumberReviews int     `gorm:"column:product_number_reviews;"`
	IsDeleted     bool    `gorm:"column:is_deleted;type: bool;default: false"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
