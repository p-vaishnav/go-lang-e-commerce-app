package tables

type Order struct {
	ID       int    `gorm:"column:order_id;primaryKey;autoIncrement;unique;not null;"`
	OrderPID string `gorm:"column:order_pid;type: varchar(50);not null;unique;"`
}

// Needs to be implemented
