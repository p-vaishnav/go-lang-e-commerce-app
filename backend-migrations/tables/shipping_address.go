package tables

import "time"

type Address struct {
	ID      int    `gorm:"column:addr_id;primaryKey;autoIncrement;unique;not null;"`
	AddrPID string `gorm:"column:addr_pid;type: varchar(50);not null;unique;"`
	UserPID string `gorm:"column:user_pid;type: varchar(50);"`
	Name    string `gorm:"column:name;type: varchar(50);"`
	// Location
	PinCode   string `gorm:"column:pincode;type: varchar(10);"`
	Phone     string `gorm:"column:phone;type: varchar(30)"`
	AltPhone  string `gorm:"column:alt_phone;type: varchar(30)"`
	Locality  string `gorm:"column:locality;type: varchar(50)"`
	Area      string `gorm:"column:area;type: varchar(50)"`
	City      string `gorm:"column:city;type: varchar(50)"`
	State     string `gorm:"column:state;type: varchar(50)"`
	LandMark  string `gorm:"column:land_mark;type: varchar(50)"`
	Type      string `gorm:"column:type;type: varchar(10)"`
	IsDeleted bool   `gorm:"column:is_deleted;type: bool;default: false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
