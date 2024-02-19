package tables

import "time"

type Review struct {
	ID         int    `gorm:"column:reviews_id;primaryKey;autoIncrement;unique;not null;"`
	ReviewPID  string `gorm:"column:reviews_pid;type: varchar(50);not null;unique;"`
	ProductPID string `gorm:"column:product_pid;type: varchar(50);not null;unique;"`
	UserPID    string `gorm:"column:user_pid;type: varchar(50);not null;unique;"`
	Content    string `gorm:"column:content; type: varchar(200)"`
	Rating     int    `gorm:"column:rating;"`
	IsDeleted  bool   `gorm:"column:is_deleted;type: bool;default: false"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
