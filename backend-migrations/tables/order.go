package tables

type Order struct {
	ID          int    `gorm:"column:order_id;primaryKey;autoIncrement;unique;not null;"`
	OrderPID    string `gorm:"column:order_pid;type: varchar(50);not null;unique;"`
	UserPID     string `gorm:"column:user_pid;type: varchar(50);"`
	AddrPid     string `gorm:"column:addr_pid;type: varchar(50);"`
	OrderStatus string `gorm:"column:order_status;type: varchar(20);"` // processed recieved cancelled
	PaymentInfo string `gorm:"column:payment_info;type: varchar(50);"`
}

// Needs to be implemented
