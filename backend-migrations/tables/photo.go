package tables

type Photo struct {
	ID       int    `gorm:"column:photo_id;primaryKey;autoIncrement;unique;not null;"`
	PhotoPID string `gorm:"column:photo_pid;type: varchar(50);not null;unique;"`
	RefPID   string `gorm:"column:ref_pid;type: varchar(50);not null;unique;"` // it's a common table to store the photo's related to user, address, review and so on
}
