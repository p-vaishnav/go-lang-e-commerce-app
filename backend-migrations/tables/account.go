package tables

type Account struct {
	ID      int    `gorm:"column:accounts_id;primaryKey;autoIncerement:true;"`
	Balance int    `gorm:"column:accounts_balance;"`
	Name    string `gorm:"column:account_name;type: varchar(10);"`
}
