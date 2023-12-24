package tables

import "time"

type User struct {
	ID        int    `gorm:"column:user_id;primaryKey;autoIncrement;unique;not null;"`
	UserPID   string `gorm:"column:user_pid;type: varchar(50);not null;unique;"`
	FirstName string `gorm:"column:user_first_name;type: varchar(20);"`
	LastName  string `gorm:"column:user_last_name;type: varchar(20)"`
	Gender    string `gorm:"column:user_gender;type: varchar(2)"`
	Email     string `gorm:"column:user_email;type: varchar(30)"`
	Phone     string `gorm:"column:user_phone;type: varchar(30)"`
	Role      string `gorm:"column:user_role;type: varchar(30)"`
	PhotoID   string `gorm:"column:user_photo;type: varchar(30)"`
	IsDeleted bool   `gorm:"column:user_is_deleted;type: bool;default: false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
