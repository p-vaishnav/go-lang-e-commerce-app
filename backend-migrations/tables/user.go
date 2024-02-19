package tables

import "time"

type User struct {
	ID        int    `gorm:"column:users_id;primaryKey;autoIncrement;unique;not null;"`
	UserPID   string `gorm:"column:users_pid;type: varchar(50);not null;unique;"`
	FirstName string `gorm:"column:users_first_name;type: varchar(20);"`
	LastName  string `gorm:"column:users_last_name;type: varchar(20)"`
	Gender    string `gorm:"column:users_gender;type: varchar(2)"`
	Email     string `gorm:"column:users_email;type: varchar(30)"`
	Phone     string `gorm:"column:users_phone;type: varchar(30)"`
	Role      string `gorm:"column:users_role;type: varchar(30)"` // SELLER, BUYER and ADMIN
	PhotoID   string `gorm:"column:users_photo;type: varchar(30)"`
	IsDeleted bool   `gorm:"column:users_is_deleted;type: bool;default: false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
